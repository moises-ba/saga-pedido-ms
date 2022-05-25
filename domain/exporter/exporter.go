package exporter

import (
	"github.com/moises-ba/saga-pedido-ms/domain/entity"
	"github.com/moises-ba/saga-pedido-ms/domain/repository"
)

type PedidoExporter interface {
	ExportStream(pedidoChan chan *entity.Pedido, storage repository.Storage, fileName string) (string, error)
	Export(pedido []*entity.Pedido) ([]byte, error)
}

type PedidoExporterFactory interface {
	Create(exportType string) (PedidoExporter, error)
}
