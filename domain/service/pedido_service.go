package service

import (
	"strings"

	"github.com/google/uuid"
	"github.com/moises-ba/saga-pedido-ms/domain"
	"github.com/moises-ba/saga-pedido-ms/domain/converter"
	"github.com/moises-ba/saga-pedido-ms/domain/dto"
	"github.com/moises-ba/saga-pedido-ms/domain/repository"
	"github.com/sirupsen/logrus"
)

func NewPedidoService(pPedidoRepository repository.PedidoRepository, pPedidoMessagingRepository repository.PedidoMessagingRepository) PedidoService {
	return &pedidoService{pedidoRepository: pPedidoRepository, pedidoMessagingRepository: pPedidoMessagingRepository}
}

type pedidoService struct {
	pedidoRepository          repository.PedidoRepository
	pedidoMessagingRepository repository.PedidoMessagingRepository
}

func (s *pedidoService) ProcessarPedido(pedidoRequest *dto.PedidoRequest) error {

	status := strings.ToUpper(pedidoRequest.Status)
	switch status {
	case "PENDING":
		pedido, err := s.CriarPedido(pedidoRequest)

		if err != nil {
			return err
		}

		logrus.Infof("Pedido %v registrado com sucesso com status %v", pedido.Id, pedido.Status)

		if pedido != nil && strings.ToUpper(pedido.Status) == "PENDING" {
			err = s.pedidoMessagingRepository.CriarPedido(converter.ConvertToPedidoEntity(pedidoRequest))
			if err != nil {
				return err
			}
			logrus.Infof("Pedido %v com status %v enviado para a fila de resposta", pedido.Id, pedido.Status)
		}

		return nil

	case "CANCELLED":
		pedido, err := s.CancelarPedido(pedidoRequest)

		if err == nil {
			logrus.Infof("Pedido %v cancelado com sucesso", pedido.Id)
		}

		return err

	case "EFFECTED":
		pedido, err := s.EfetivarPedido(pedidoRequest)

		if err == nil {
			logrus.Infof("Pedido %v efetivado com sucesso", pedido.Id)
		}

		return err

	default:
		logrus.Warnf("Ação %v para o pedido %v não conhecida, ignorando mensagem", status, pedidoRequest.Id)
		return nil
	}
}

func (s *pedidoService) ListarPedidos(userId string) ([]*dto.PedidoResponse, error) {
	return nil, nil
}

func (s *pedidoService) CriarPedido(pedido *dto.PedidoRequest) (*dto.PedidoResponse, error) {

	pedidoEntity := converter.ConvertToPedidoEntity(pedido)
	if strings.Trim(pedidoEntity.Id, "") == "" {
		pedidoEntity.Id = uuid.NewString() //cria um id para o pedido
	}
	pedidoEntity.Status = "PENDING"

	pedidoEntity, err := s.pedidoRepository.CriarPedido(pedidoEntity)

	var pedidoResponse *dto.PedidoResponse = nil
	if err == domain.ErrDuplicateKey {
		logrus.Warnf("Pedido %v ja processado, obtendo o status ", pedido.Id)
		err = nil
		pedidoResponse, err = s.ObterPedido(pedido.User.Id, pedido.Id)
		if err != nil {
			return nil, err
		}

		if pedidoResponse != nil {
			logrus.Infof("Pedido %v possui status ", pedidoResponse.Status)
		}
	} else {
		pedidoResponse = converter.ConvertToPedidoResponse(pedidoEntity)
	}

	//envia para a fila de pedido response
	err = s.pedidoMessagingRepository.CriarPedido(pedidoEntity)
	if err != nil {
		pedidoResponse = nil
	}

	return pedidoResponse, err
}

func (s *pedidoService) ObterPedido(userId, pPedidoId string) (*dto.PedidoResponse, error) {
	pedidoEntity, err := s.pedidoRepository.ObterPedido(userId, pPedidoId)

	if err != nil {
		if err == domain.ErrRecordNofFound {
			err = nil
		}
	}

	return converter.ConvertToPedidoResponse(pedidoEntity), nil
}

func (s *pedidoService) CancelarPedido(pedido *dto.PedidoRequest) (*dto.PedidoResponse, error) {
	return nil, nil
}

func (s *pedidoService) EfetivarPedido(pedido *dto.PedidoRequest) (*dto.PedidoResponse, error) {
	return nil, nil
}
