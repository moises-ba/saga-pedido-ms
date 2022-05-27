package exporter

import (
	"github.com/moises-ba/saga-pedido-ms/domain/entity"
	"github.com/moises-ba/saga-pedido-ms/domain/repository"
)

type PedidoExporter interface {
	Export(pedido []*entity.Pedido) ([]byte, error)
}

type PedidoExporterStream interface {
	PedidoExporter
	ExportStream(pedidoChan chan *entity.Pedido, storageStream repository.StorageStream, fileName string) (string, error)
}

type PedidoExporterFactory interface {
	Create(exportType string) (PedidoExporter, error)
	CreateExporterStream(exportType string) (PedidoExporterStream, error)
}
