package handler

import (
	"encoding/json"

	"github.com/moises-ba/saga-pedido-ms/domain/dto"
	"github.com/moises-ba/saga-pedido-ms/domain/service"
	"github.com/sirupsen/logrus"
)

type KakfkaController interface {
	OnPedidoCommandResquest() func(pMessage string) error
}

type kakfkaController struct {
	pedidoService service.PedidoService
}

func NewKafkaController(pedidoService service.PedidoService) *kakfkaController {
	return &kakfkaController{pedidoService: pedidoService}
}

func (controller *kakfkaController) OnPedidoCommandResquest(pMessage string) error {
	var err error
	pedidoRequest := &dto.PedidoRequest{}
	if err = json.Unmarshal([]byte(pMessage), &pedidoRequest); err != nil {
		logrus.Warnf("Mensagem do pedido no formato inválido: %v", pMessage)
		return nil
	}

	err = controller.pedidoService.ProcessarPedido(pedidoRequest)
	if err != nil {
		logrus.Errorf("Erro ao processar o pedido: %v", err.Error())
		return err
	}

	return nil
}
