package service

import (
	"Order-Matching-System/models"
	"container/heap"
)

type OrderHeap struct {
	orders []*models.Order
	side   string
}

func NewOrderHeap(side string) *OrderHeap {
	oh := &OrderHeap{side: side}
	heap.Init(oh)
	return oh
}

func (oh OrderHeap) Len() int { return len(oh.orders) }

func (oh OrderHeap) Less(i, j int) bool {
	if oh.side == "buy" {
		return oh.orders[i].Price > oh.orders[j].Price
	}
	return oh.orders[i].Price < oh.orders[j].Price
}

func (oh OrderHeap) Swap(i, j int) {
	oh.orders[i], oh.orders[j] = oh.orders[j], oh.orders[i]
}

func (oh *OrderHeap) Push(x interface{}) {
	oh.orders = append(oh.orders, x.(*models.Order))
}

func (oh *OrderHeap) Pop() interface{} {
	n := len(oh.orders)
	item := oh.orders[n-1]
	oh.orders = oh.orders[:n-1]
	return item
}

func (oh *OrderHeap) Top() *models.Order {
	if len(oh.orders) == 0 {
		return nil
	}
	return oh.orders[0]
}
