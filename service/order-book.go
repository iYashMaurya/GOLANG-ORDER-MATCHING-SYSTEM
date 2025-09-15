package service

import "Order-Matching-System/models"

type OrderBook struct {
	Symbol string
	Bids   *OrderHeap
	Asks   *OrderHeap
}

func NewOrderBook(symbol string) *OrderBook {
	return &OrderBook{
		Symbol: symbol,
		Bids:   NewOrderHeap("buy"),
		Asks:   NewOrderHeap("sell"),
	}
}

func (ob *OrderBook) AddOrder(order *models.Order) {
	if order.Side == "buy" {
		ob.Bids.Push(order)
	} else {
		ob.Asks.Push(order)
	}
}
