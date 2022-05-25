package messagebroker

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type KafkaMessageBroker interface {
	Consume(pTopic string, numThreads int, handle func(pMessage string) error)
	CreateWriter(topic string) *kafka.Writer
	CloseReaders()
}

type kafkaMessageBroker struct {
	brokers         []string
	groupId         string
	consumers       map[string]*kafka.Reader
	context         context.Context
	doneReadMessage func()
	stopConsumers   bool           //flag q indica se o consumer deve finalizar
	wgWaitClose     sync.WaitGroup //wait group responsavel por gerenciar a espera do fechamento dos consumers
}

func NewKafkaBroker(pBrokers []string, pGroupId string) KafkaMessageBroker {

	ctx, cancel := context.WithCancel(context.Background())
	return &kafkaMessageBroker{brokers: pBrokers,
		groupId:         pGroupId,
		consumers:       make(map[string]*kafka.Reader),
		stopConsumers:   false,
		context:         ctx,
		doneReadMessage: cancel}
}

func (k *kafkaMessageBroker) CreateWriter(topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: k.brokers,
		Topic:   topic,
	})
}

func (k *kafkaMessageBroker) Consume(pTopic string, numTrheads int, handle func(pMessage string) error) {
	k.stopConsumers = false
	for i := 0; i < numTrheads; i++ {
		k.wgWaitClose.Add(1)
		go func() {

			readerConfig := kafka.ReaderConfig{
				Brokers: k.brokers,
				Topic:   pTopic,
				GroupID: k.groupId,
			}
			r := kafka.NewReader(readerConfig)
			defer func() { //funcao para fechar o consumidor
				errClose := r.Close()
				if errClose != nil {
					logrus.Errorf("Falha o fechar o consumidor: %v", errClose.Error())
				}
				k.wgWaitClose.Done()
			}()

			kafkaReaderId, _ := uuid.NewRandom()
			k.consumers[kafkaReaderId.String()] = r

			for !k.stopConsumers {

				msg, err := r.FetchMessage(k.context)
				if err != nil {
					if err == context.Canceled {
						log.Info("finalizando consumer...")
					} else {
						log.Errorf("Não foi possível ler da fila %v, erro type: %T error: %v, ", r.Config().Topic, err, err.Error())
					}

					continue
				}

				if err = handle(string(msg.Value)); err == nil {
					if err := r.CommitMessages(k.context, msg); err != nil {
						log.Errorf("Falha ao commitar no kafka:", err)
					}
				} else {
					log.Error("Falha ao processar pedido, reiniciando o consumer. erro -> ", err)
					r.Close()
					r = kafka.NewReader(readerConfig)
					k.consumers[kafkaReaderId.String()] = r
				}

			}

		}()

	}

}

func (k *kafkaMessageBroker) CloseReaders() {
	log.Infof("Fechando %v consumer(s)...", len(k.consumers))
	k.stopConsumers = true
	if k.doneReadMessage != nil {
		k.doneReadMessage() //interrompe do FetchMessage do kafka
	}
	k.wgWaitClose.Wait()
	logrus.Info("Consumers fechados")
}
