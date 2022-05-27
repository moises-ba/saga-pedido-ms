package exporterprovider

import (
	"github.com/moises-ba/saga-pedido-ms/domain/entity"

	"github.com/moises-ba/saga-pedido-ms/domain/exporter"
)

type pedidoPDFExporter struct {
}

func NewPedidoPDFExporter() exporter.PedidoExporter {
	return &pedidoPDFExporter{}
}

func (e *pedidoPDFExporter) Export(pedido []*entity.Pedido) ([]byte, error) {
	return nil, nil
}
