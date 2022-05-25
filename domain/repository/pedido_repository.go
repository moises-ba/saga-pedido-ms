package repository

import "github.com/moises-ba/saga-pedido-ms/domain/entity"

type PedidoReader interface {
	ListarPedidos(userId string) ([]*entity.Pedido, error)
	ObterPedido(userId string, pPedidoId string) (*entity.Pedido, error)
}

type PedidoWriter interface {
	CriarPedido(pedido *entity.Pedido) (*entity.Pedido, error)
}

type PedidoRepository interface {
	PedidoReader
	PedidoWriter
}
