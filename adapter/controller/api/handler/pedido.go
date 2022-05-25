package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moises-ba/saga-pedido-ms/adapter/presenter"
	"github.com/moises-ba/saga-pedido-ms/domain/dto"
	"github.com/moises-ba/saga-pedido-ms/domain/service"
	"golang.org/x/mod/sumdb/storage"
)

func CriarPedido(pedidoService service.PedidoService) func(c *gin.Context) {

	return func(c *gin.Context) {
		usuarioHeader := c.GetHeader("usuarioId")
		if usuarioHeader == "" {
			c.JSON(400, gin.H{"message": "parametro usuarioId no header é requerido"})
			return
		}

		pedidoRequest := &dto.PedidoRequest{}

		if err := c.BindJSON(pedidoRequest); err != nil {
			c.JSON(400, gin.H{"message": fmt.Sprintf("Falha ao criar pedido: %v", err.Error())})
			return
		}

		pedido, err := pedidoService.CriarPedido(pedidoRequest)
		if err != nil {
			c.JSON(400, gin.H{"message": fmt.Sprintf("Falha ao Criar pedido: %v", err.Error())})
			return
		}

		c.JSON(200, dto.PedidoCreatedResponse{Id: pedido.Id})

	}

}

func ListarPedidos(pedidoService service.PedidoService) func(c *gin.Context) {

	return func(c *gin.Context) {
		usuarioHeader := c.GetHeader("usuarioId")
		if usuarioHeader == "" {
			c.JSON(400, gin.H{"message": "parametro usuarioId no header é requerido"})
			return
		}

		pedidos, err := pedidoService.ListarPedidos("fulano")
		if err != nil {
			c.JSON(500, gin.H{"message": fmt.Sprintf("Falha ao listar pedidos: %v", err.Error())})
			return
		}

		c.JSON(200, presenter.NewPedidoRestPresenterList(pedidos))

	}

}

func ObterPedido(pedidoService service.PedidoService) func(c *gin.Context) {

	return func(c *gin.Context) {

		usuarioHeader := c.GetHeader("usuarioId")
		if usuarioHeader == "" {
			c.JSON(400, gin.H{"message": "parametro usuarioId no header é requerido"})
			return
		}

		paramPedidoId := c.Param("pedidoId")
		if paramPedidoId == "" {
			c.JSON(400, gin.H{"message": "parametro predidoId é requerido"})
			return
		}
		pedido, err := pedidoService.ObterPedido("fulano", paramPedidoId)
		if err != nil {
			c.JSON(500, gin.H{"message": fmt.Sprintf("Falha ao listar pedidos: %v", err.Error())})
			return
		}

		if pedido == nil {
			c.JSON(400, gin.H{"message": fmt.Sprintf("Pedido %v não encontrado", paramPedidoId)})
			return
		}

		c.JSON(200, presenter.NewPedidoRestPresenter(pedido))

	}

}

func ExportarPedidosAsync(service service.PedidoExportService, storage storage.Storage) func(c *gin.Context) {

	return func(c *gin.Context) {
		usuarioHeader := c.GetHeader("usuarioId")
		if usuarioHeader == "" {
			c.JSON(400, gin.H{"message": "parametro usuarioId no header é requerido"})
			return
		}

		exportType := c.GetHeader("export_type")
		if exportType == "" {
			exportType = "csv" //default
		}

		urlFile, err := service.ExportarPedidosAsync(usuarioHeader, exportType, storage)
		if err != nil {
			c.JSON(500, gin.H{"message": fmt.Sprintf("Falha ao listar pedidos: %v", err.Error())})
			return
		}

		c.JSON(200, gin.H{"url": urlFile})

	}

}

func ExportarPedidos(service service.PedidoExportService, storage storage.Storage) func(c *gin.Context) {

	return func(c *gin.Context) {
		usuarioHeader := c.GetHeader("usuarioId")
		if usuarioHeader == "" {
			c.JSON(400, gin.H{"message": "parametro usuarioId no header é requerido"})
			return
		}

		exportType := c.GetHeader("export_type")
		if exportType == "" {
			exportType = "csv" //default
		}

		bytesContent, err := service.ExportarPedidos(usuarioHeader, exportType)
		if err != nil {
			c.JSON(500, gin.H{"message": fmt.Sprintf("Falha ao listar pedidos: %v", err.Error())})
			return
		}

		c.Writer.Write(bytesContent)

	}

}
