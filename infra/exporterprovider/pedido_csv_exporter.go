package exporterprovider

import (
	"fmt"

	"github.com/moises-ba/saga-pedido-ms/domain/entity"
	"github.com/moises-ba/saga-pedido-ms/domain/repository"

	"github.com/moises-ba/saga-pedido-ms/domain/exporter"
)

type pedidoCSVExporter struct {
}

func NewPedidoCSVExporter() exporter.PedidoExporterStream {
	return &pedidoCSVExporter{}
}

func (e *pedidoCSVExporter) Export(pedidos []*entity.Pedido) ([]byte, error) {

	csvContent := createCSVHeader()
	for _, pedido := range pedidos {
		csvContent += createCSVRow(pedido)
	}

	return []byte(csvContent), nil
}

func (e *pedidoCSVExporter) ExportStream(pedidoChan chan *entity.Pedido, storageStream repository.StorageStream, fileName string) (string, error) {

	chanContentByte := make(chan []byte)
	go func() {
		defer close(chanContentByte)

		chanContentByte <- []byte(createCSVHeader())
		for pedido := range pedidoChan {
			chanContentByte <- []byte(createCSVRow(pedido))
		}
	}()

	storageStream.StoreStream(chanContentByte, fileName)

	return "", nil
}

func createCSVHeader() string {
	return "ID_PEDIDO;ID_USUARIO;USUARIO;DATA_PEDIDO;STATUS;\n"

}

func createCSVRow(pedido *entity.Pedido) string {
	return fmt.Sprintf("%v;%v;%v;%v;%v;\n", pedido.Id, pedido.User.Id, pedido.User.Email, pedido.Date.String(), pedido.Status)
}
