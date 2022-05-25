package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moises-ba/saga-pedido-ms/domain/service"
	"github.com/moises-ba/saga-pedido-ms/infra/db"
	"github.com/moises-ba/saga-pedido-ms/infra/exporterprovider"
	"github.com/moises-ba/saga-pedido-ms/infra/messagebroker"
	"github.com/moises-ba/saga-pedido-ms/infra/storageprovider"

	"github.com/moises-ba/saga-pedido-ms/adapter/controller/api/handler"
	messaginHandler "github.com/moises-ba/saga-pedido-ms/adapter/controller/messaging/handler"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	godotenv.Load() //inicia a leitura das variaveis de ambiente no arquivo .env

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		log.Infof("Recebido sinal de finalização, encerrando... ", sig.String())
		done <- true
	}()

	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.SetFormatter(formatter)

	//cria conexao com o mongo
	mongoDb := connectMongo()
	//criando repository
	pedidoRepository := db.NewPedidoMongoRepository(mongoDb)

	//configurando o kafka
	messageBroker := connectToKafka(os.Getenv("KAFKA_BROKERS"))
	topic_request := "pedido_request"
	topic_response := "pedido_response"
	pedidoKafkaWriter := messageBroker.CreateWriter(topic_response)
	pedidoMessagingRepository := messagebroker.NewPedidoKafkaMessagingRepository(pedidoKafkaWriter)

	//criando o servico de pedido
	pedidoService := service.NewPedidoService(pedidoRepository, pedidoMessagingRepository)

	//servico de exportacao
	pedidoExporterFactory := exporterprovider.NewPedidoExporterFactory()
	exporterStorage := storageprovider.NewLocalFileStorage("/tmp/")
	pedidoExporterService := service.NewPedidoExportService(pedidoRepository, pedidoExporterFactory)

	//consome os comandos da fila kafka
	numThreadsConsumerPedido, err := strconv.Atoi(os.Getenv("KAFKA_PEDIDO_CONSUMER_NUM_THREADS"))
	if err != nil {
		log.Fatalf("valor da variavel de numero de consumers deve ser um inteiro")
	}

	kafkaController := messaginHandler.NewKafkaController(pedidoService)
	go messageBroker.Consume(topic_request, numThreadsConsumerPedido, kafkaController.OnPedidoCommandResquest)

	//iniciando o restAPIServer
	r := gin.Default()

	r.POST("/pedido/", handler.CriarPedido(pedidoService))
	r.GET("/pedidos", handler.ListarPedidos(pedidoService))
	r.GET("/pedido/:pedidoId", handler.ObterPedido(pedidoService))
	r.GET("/pedido/exportar/", handler.ExportarPedidos(pedidoExporterService, exporterStorage))

	go func() { //iniciando o servidor gin gonic em outra thread
		err := r.Run(":" + os.Getenv("PEDIDO_SERVER_PORT"))
		if err != nil {
			log.Fatal("Falha ao iniciar o Gin-gonnic server: %v", err)
		}
	}()

	<-done
	messageBroker.CloseReaders()

}

func connectMongo() *mongo.Database {
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL")).SetAuth(options.Credential{
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Falha ao conectar no mongo", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Falha ao efetuar o ping no mongo", err)
	}

	return client.Database(os.Getenv("MONGO_DATABASE"))

}

func connectToKafka(kafkaBrokersAddress string) messagebroker.KafkaMessageBroker {

	log.Infof("kafka_brokers: %v", kafkaBrokersAddress)

	log.Debug("verificando se o kafka esta ativo")
	_, err := kafka.Dial("tcp", kafkaBrokersAddress)
	if err != nil {
		log.Fatal("Falha ao conectar no kafka: ", err.Error())
	}

	return messagebroker.NewKafkaBroker(strings.Split(kafkaBrokersAddress, ","), "grp_pedido")
}
