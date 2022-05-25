package service

import (
	"fmt"
	"time"

	"github.com/moises-ba/saga-pedido-ms/domain/entity"
	"github.com/moises-ba/saga-pedido-ms/domain/exporter"
	"github.com/moises-ba/saga-pedido-ms/domain/repository"
	"github.com/sirupsen/logrus"
)

type pedidoExportService struct {
	pedidoRepository      repository.PedidoRepository
	pedidoExporterFactory exporter.PedidoExporterFactory
}

func NewPedidoExportService(pedidoRepository repository.PedidoRepository,
	pedidoExporterFactory exporter.PedidoExporterFactory) PedidoExportService {

	return &pedidoExportService{
		pedidoRepository:      pedidoRepository,
		pedidoExporterFactory: pedidoExporterFactory,
	}
}

func (p *pedidoExportService) ExportarPedidosAsync(userId, exportType string, storage repository.Storage) (string, error) {

	var pedidos []*entity.Pedido
	var err error
	var fullPath string
	var pedidoExporter exporter.PedidoExporter
	var reportPedidosByte []byte

	if pedidos, err = p.pedidoRepository.ListarPedidos(userId); err != nil {
		return "", fmt.Errorf("Falha ao listar pedidos: %v", err.Error())
	}

	if pedidoExporter, err = p.pedidoExporterFactory.Create(exportType); err != nil {
		return "", fmt.Errorf("Falha ao criar exporter: %v", err.Error())
	}

	if reportPedidosByte, err = pedidoExporter.Export(pedidos); err != nil {
		logrus.Error("falha ao exportar pedidos")
		return "", err
	}

	fileName := fmt.Sprintf("%v_%v", userId, time.Now().String())

	if err = storage.Store(reportPedidosByte, fileName); err != nil {
		logrus.Error("falha ao gravar pedido exportado")
		return "", err
	}

	return fullPath, nil
}

func (p *pedidoExportService) ExportarPedidos(userId, exportType string) ([]byte, error) {

	var pedidos []*entity.Pedido
	var err error
	var pedidoExporter exporter.PedidoExporter
	var reportPedidosByte []byte

	if pedidos, err = p.pedidoRepository.ListarPedidos(userId); err != nil {
		return nil, fmt.Errorf("Falha ao listar pedidos: %v", err.Error())
	}

	if pedidoExporter, err = p.pedidoExporterFactory.Create(exportType); err != nil {
		return nil, fmt.Errorf("Falha ao criar exporter: %v", err.Error())
	}

	if reportPedidosByte, err = pedidoExporter.Export(pedidos); err != nil {
		logrus.Error("falha ao exportar pedidos")
		return nil, err
	}

	return reportPedidosByte, nil
}
