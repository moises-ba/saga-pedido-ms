package service

import (
	"github.com/moises-ba/saga-pedido-ms/domain/dto"
	"github.com/moises-ba/saga-pedido-ms/domain/repository"
)

type PedidoService interface {
	ListarPedidos(userId string) ([]*dto.PedidoResponse, error)
	ProcessarPedido(pedidoRequest *dto.PedidoRequest) error
	CriarPedido(pedido *dto.PedidoRequest) (*dto.PedidoResponse, error)
	ObterPedido(userId, pPedidoId string) (*dto.PedidoResponse, error)
	CancelarPedido(pedido *dto.PedidoRequest) (*dto.PedidoResponse, error)
	EfetivarPedido(pedido *dto.PedidoRequest) (*dto.PedidoResponse, error)
}

type PedidoExportService interface {
	ExportarPedidosAsync(userId, exportType string, storage repository.Storage) (string, error)
	ExportarPedidos(userId, exportType string) ([]byte, error)
}
