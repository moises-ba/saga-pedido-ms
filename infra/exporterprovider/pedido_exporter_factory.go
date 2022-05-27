package exporterprovider

import (
	"fmt"

	"github.com/moises-ba/saga-pedido-ms/domain/exporter"
)

type pedidoExporterFactory struct{}

func NewPedidoExporterFactory() exporter.PedidoExporterFactory {
	return &pedidoExporterFactory{}
}

func (f *pedidoExporterFactory) Create(exportType string) (exporter.PedidoExporter, error) {
	switch exportType {
	case "pdf":
		return NewPedidoPDFExporter(), nil
	case "xlsx":
		return NewPedidoXLSXExporter(), nil
	case "csv":
		return NewPedidoCSVExporter(), nil
	default:
		return nil, fmt.Errorf("exportador %v inexistente", exportType)

	}
}
func (f *pedidoExporterFactory) CreateExporterStream(exportType string) (exporter.PedidoExporterStream, error) {
	switch exportType {
	case "csv":
		return NewPedidoCSVExporter(), nil
	default:
		return nil, fmt.Errorf("exportador %v inexistente", exportType)

	}
}
