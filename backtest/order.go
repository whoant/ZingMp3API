package backtest

import (
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
)

type OrderType int
type OrderState int

const (
	BID OrderType = iota
	ASK
)

var OrderTypes = [2]string{
	"BID",
	"ASK",
}

func (o *OrderType) String() string {
	return OrderTypes[*o]
}

func (o *OrderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(*o))
}

func (o *OrderType) UnmarshalJSON(data []byte) error {
	var i int
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}
	*o = OrderType(i)
	return nil
}

const (
	ENABLED OrderState = iota
	CANCEL
	DONE
)

var OrderStates = [3]string{
	"ENABLED",
	"CANCEL",
	"DONE",
}

func (o *OrderState) String() string {
	return OrderStates[*o]
}

func (o *OrderState) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(*o))
}

func (o *OrderState) UnmarshalJSON(data []byte) error {
	var i int
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}
	*o = OrderState(i)
	return nil
}

type Order struct {
	OrderType        OrderType  `json:"orderType"`
	State            OrderState `json:"state"`
	OpenedAt         time.Time  `json:"openedAt"`
	CanceledAt       time.Time  `json:"canceledAt"`
	DoneAt           time.Time  `json:"doneAt"`
	Amount           float64    `json:"amount"`
	OpenPrice        float64    `json:"openPrice"`
	TakeProfitPrice  float64    `json:"takeProfitPrice"`
	CancelOrderPrice float64    `json:"cancelOrderPrice"`
}

func NewOrder(openingOrder *OpeningOrder, amount float64, openPrice float64, openedAt time.Time) *Order {
	return &Order{
		OpenedAt:         openedAt,
		OrderType:        openingOrder.OrderType,
		State:            ENABLED,
		Amount:           amount,
		OpenPrice:        openPrice,
		TakeProfitPrice:  openingOrder.TakeProfitPrice,
		CancelOrderPrice: openingOrder.CancelOrderPrice,
	}
}

func (order *Order) IsEnable() bool {
	return order.State == ENABLED
}

func (order *Order) IsDone() bool {
	return order.State == DONE
}

func (order *Order) IsCancel() bool {
	return order.State == CANCEL
}

func (order *Order) MarkMatched(doneAt time.Time) {
	order.State = DONE
	order.DoneAt = doneAt
}

func (order *Order) MarkCancel(cancelAt time.Time) {
	order.State = CANCEL
	order.CanceledAt = cancelAt
}

func (order *Order) Log() {
	log.Info().Str("type", "history_order").
		Str("order_type", parseOrderType2Str(order.OrderType)).
		Str("state", parseOrderState2Str(order.State)).
		Time("open_at", order.OpenedAt).
		Time("cancel_at", order.CanceledAt).
		Time("done_at", order.DoneAt).
		Float64("amount", order.Amount).
		Float64("open_price", order.OpenPrice).
		Float64("take_profit_price", order.TakeProfitPrice).
		Float64("cancel_order_price", order.CancelOrderPrice).
		Msg("query history order")
}

func parseOrderType2Str(orderType OrderType) string {
	return OrderTypes[orderType]
}

func parseOrderState2Str(state OrderState) string {
	return OrderStates[state]
}
