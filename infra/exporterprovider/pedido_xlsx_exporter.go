package exporterprovider

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/moises-ba/saga-pedido-ms/domain/entity"
	"github.com/moises-ba/saga-pedido-ms/domain/exporter"
	"github.com/moises-ba/saga-pedido-ms/domain/repository"
	"github.com/tealeg/xlsx"
)

type pedidoXLSXExporter struct {
}

func NewPedidoXLSXExporter() exporter.PedidoExporter {
	return &pedidoXLSXExporter{}
}

func (e *pedidoXLSXExporter) Export(pedidos []*entity.Pedido) ([]byte, error) {

	file, sheet, err := createXlsFileWithHeader()
	if err != nil {
		return nil, err
	}

	for _, pedido := range pedidos {
		createRow(sheet, pedido)
	}

	var buf bytes.Buffer
	err = file.Write(&buf)
	if err != nil {
		return nil, fmt.Errorf("falha ao gravar no buffer: %v", err.Error())
	}

	return buf.Bytes(), nil
}

func (e *pedidoXLSXExporter) ExportStream(pedidoChan chan *entity.Pedido, storage repository.Storage, fileName string) (string, error) {
	return "", errors.New("Exportacao via stream nao suportado")
}

func createXlsFileWithHeader() (*xlsx.File, *xlsx.Sheet, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("pedidos")
	if err != nil {
		return nil, nil, fmt.Errorf("falha ao criar sheet na planilha: %v", err.Error())
	}

	row := sheet.AddRow() //cabecalho
	row.AddCell().Value = "ID_PEDIDO"
	row.AddCell().Value = "ID_USUARIO"
	row.AddCell().Value = "USUARIO"
	row.AddCell().Value = "DATA_PEDIDO"
	row.AddCell().Value = "STATUS"

	return file, sheet, nil
}

func createRow(sheet *xlsx.Sheet, pedido *entity.Pedido) {
	panic("unimplemented")
}
