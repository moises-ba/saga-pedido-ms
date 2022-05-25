package presenter

import (
	"time"

	"github.com/moises-ba/saga-pedido-ms/domain/dto"
)

type pedidoRestPresenter struct {
	Id            string                  `json:"id"`
	User          *usuarioPresenter       `json:"user"`
	Items         []*itemPresenter        `json:"items"`
	Status        string                  `json:"status"`
	Reason        string                  `json:"reason"`
	Date          *time.Time              `json:"date"`
	PaymentDetail *paymentDetailPresenter `json:"paymentDetail"`
	TotalPrice    float64                 `json:"totalPrice"`
}

type usuarioPresenter struct {
	Id string `json:"id"`
}

type itemPresenter struct {
	Product   *productPresenter `json:"product"`
	Quantity  int32             `json:"quantity"`
	UnitPrice float64           `json:"unitPrice"`
	Total     float64           `json:"total"` //preco calculado Quantity * UnitPrice
}

type paymentDetailPresenter struct {
	CardType   string `json:"cardType"`   //VISA etc
	CardNumber string `json:"cardNumber"` //000000
}

type productPresenter struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unitPrice"`
}

func NewPedidoRestPresenter(pedidoResponse *dto.PedidoResponse) *pedidoRestPresenter {
	return &pedidoRestPresenter{}
}

func NewPedidoRestPresenterList(pedidosResponse []*dto.PedidoResponse) []*pedidoRestPresenter {
	pedidosPresenters := make([]*pedidoRestPresenter, len(pedidosResponse))
	for i, pedidoResponse := range pedidosResponse {
		pedidosPresenters[i] = NewPedidoRestPresenter(pedidoResponse)
	}
	return pedidosPresenters
}
