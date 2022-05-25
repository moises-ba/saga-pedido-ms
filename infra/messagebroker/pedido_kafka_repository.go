package messagebroker

import (
	"context"
	"encoding/json"

	"github.com/moises-ba/saga-pedido-ms/domain/entity"
	"github.com/moises-ba/saga-pedido-ms/domain/repository"
	"github.com/segmentio/kafka-go"
)

type pedidoKafkaMessagingRepository struct {
	writer *kafka.Writer
}

func NewPedidoKafkaMessagingRepository(writer *kafka.Writer) repository.PedidoMessagingRepository {
	return &pedidoKafkaMessagingRepository{writer: writer}
}

func (p *pedidoKafkaMessagingRepository) CriarPedido(pedido *entity.Pedido) error {

	jsonMessage, err := json.Marshal(pedido)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(context.TODO(), kafka.Message{
		Key:   []byte(pedido.User.Id),
		Value: jsonMessage,
	})

}
