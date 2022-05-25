package converter

import (
	"github.com/moises-ba/saga-pedido-ms/domain/dto"
	"github.com/moises-ba/saga-pedido-ms/domain/entity"
)

//converte um model de pedido em uma entidade pedido
func ConvertToPedidoEntity(pPedido *dto.PedidoRequest) *entity.Pedido {

	pedidoEntity := &entity.Pedido{
		Id: pPedido.Id,
		User: &entity.User{
			Id: pPedido.User.Id,
		},
		Status: pPedido.Status,
		Reason: pPedido.Reason,
		PaymentDetail: &entity.PaymentDetail{
			CardType:   pPedido.PaymentDetail.CardType,
			CardNumber: pPedido.PaymentDetail.CardNumber,
		},
	}

	var totalPedido float64 = 0
	if len(pPedido.Items) > 0 {
		pedidoEntity.Items = make([]*entity.Item, 0)
		for _, pItem := range pPedido.Items {
			itemEntity := &entity.Item{
				Quantity:   pItem.Quantity,
				TotalPrice: float64(pItem.Quantity) * pItem.Product.UnitPrice,
			}
			totalPedido += itemEntity.TotalPrice //atualiza o valor total do pedido
			pedidoEntity.Items = append(pedidoEntity.Items, itemEntity)

		}
	}

	pedidoEntity.TotalPrice = totalPedido //atribui o valor total do pedido

	return pedidoEntity

}

//converte uma entidade de pedido em um objetos de resposta para a fila
func ConvertToPedidoResponse(pPedido *entity.Pedido) *dto.PedidoResponse {

	pedidoResponse := &dto.PedidoResponse{
		Id:     pPedido.Id,
		User:   &dto.User{Id: pPedido.User.Id},
		Status: pPedido.Status,
		Reason: pPedido.Reason,
		PaymentDetail: &dto.PaymentDetail{
			CardType:   pPedido.PaymentDetail.CardType,
			CardNumber: pPedido.PaymentDetail.CardNumber,
		},
		TotalPrice: pPedido.TotalPrice,
		Items:      make([]*dto.Item, 0),
	}

	for _, v := range pPedido.Items {
		pedidoResponse.Items = append(pedidoResponse.Items, &dto.Item{
			Product: &dto.Product{
				Id:        v.Product.Id,
				Name:      v.Product.Name,
				UnitPrice: v.Product.UnitValue,
			},
			Total:    v.TotalPrice,
			Quantity: v.Quantity,
		})
	}

	return pedidoResponse
}
