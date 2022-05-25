package repository

import (
	"github.com/moises-ba/saga-pedido-ms/domain/entity"
)

type PedidoMessagingRepository interface {
	CriarPedido(pedido *entity.Pedido) error
}
