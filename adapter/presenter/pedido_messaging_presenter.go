package presenter

import (
	"time"

	"github.com/moises-ba/saga-pedido-ms/domain/dto"
)

type pedidoMessagingPresenter struct {
	Id            string                           `json:"id"`
	User          *usuarioMessagingPresenter       `json:"user"`
	Items         []*itemMessagingPresenter        `json:"items"`
	Status        string                           `json:"status"`
	Reason        string                           `json:"reason"`
	Date          *time.Time                       `json:"date"`
	PaymentDetail *paymentDetailMessagingPresenter `json:"paymentDetail"`
	TotalPrice    float64                          `json:"totalPrice"`
}

type usuarioMessagingPresenter struct {
	Id string `json:"id"`
}

type itemMessagingPresenter struct {
	Product   *productMessagingPresenter `json:"product"`
	Quantity  int32                      `json:"quantity"`
	UnitPrice float64                    `json:"unitPrice"`
	Total     float64                    `json:"total"` //preco calculado Quantity * UnitPrice
}

type paymentDetailMessagingPresenter struct {
	CardType   string  `json:"cardType"`   //VISA etc
	CardNumber string  `json:"cardNumber"` //000000
	TotalPrice float64 `json:"totalPrice"`
}

type productMessagingPresenter struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unitPrice"`
}

func NewPedidoMessagingPresenter(pedidoResponse *dto.PedidoResponse) *pedidoMessagingPresenter {
	return &pedidoMessagingPresenter{
		Id: pedidoResponse.Id,
		User: &usuarioMessagingPresenter{
			Id: pedidoResponse.User.Id,
		},
		PaymentDetail: &paymentDetailMessagingPresenter{
			CardType:   pedidoResponse.PaymentDetail.CardType,
			CardNumber: pedidoResponse.PaymentDetail.CardNumber,
			TotalPrice: pedidoResponse.TotalPrice,
		},
	}
}
