package backtest

type ExchangeHandler struct {
	fee           float64
	HistoryOrders []*Order
}

func NewExchangeHandler(fee float64) *ExchangeHandler {
	return &ExchangeHandler{
		fee:           fee,
		HistoryOrders: make([]*Order, 0),
	}
}

func (handler *ExchangeHandler) CountOpenOrder() int {
	count := 0
	for _, order := range handler.HistoryOrders {
		if order.State == ENABLED {
			count++
		}
	}

	return count
}

func (handler *ExchangeHandler) MatchingOrder(price DataPoint) {
	for _, currentOrder := range handler.HistoryOrders {
		if !currentOrder.IsEnable() {
			continue
		}

		if currentOrder.TakeProfitPrice >= price.LowPrice() && currentOrder.TakeProfitPrice <= price.HighPrice() {
			currentOrder.MarkMatched(price.Timestamp())
		}
	}
}

func (handler *ExchangeHandler) CancelOrder(price DataPoint) {
	for _, currentOrder := range handler.HistoryOrders {
		if !currentOrder.IsEnable() {
			continue
		}

		if currentOrder.CancelOrderPrice >= price.LowPrice() && currentOrder.CancelOrderPrice <= price.HighPrice() {
			currentOrder.MarkCancel(price.Timestamp())
		}
	}
}
