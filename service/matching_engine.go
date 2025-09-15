package service

import (
	"Order-Matching-System/models"
	"database/sql"
	"log"
)

type MatchingEngine struct {
	db        *sql.DB
	orderBook map[string]*OrderBook
}

func NewMatchingEngine(db *sql.DB) *MatchingEngine {
	return &MatchingEngine{
		db:        db,
		orderBook: make(map[string]*OrderBook),
	}
}

func (me *MatchingEngine) LoadOpenOrders() error {
	rows, err := me.db.Query("SELECT id, symbol, side, type, price, quantity, status, created_at FROM orders WHERE status IN ('open','partially_filled')")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.Symbol, &o.Side, &o.Type, &o.Price, &o.Quantity, &o.Status, &o.CreatedAt); err != nil {
			return err
		}

		if _, exists := me.orderBook[o.Symbol]; !exists {
			me.orderBook[o.Symbol] = NewOrderBook(o.Symbol)
		}
		me.orderBook[o.Symbol].AddOrder(&o)
	}
	return nil
}

func (me *MatchingEngine) PlaceOrder(o *models.Order) error {
	res, err := me.db.Exec(
		"INSERT INTO orders (symbol, side, type, price, quantity, status) VALUES (?, ?, ?, ?, ?, ?)",
		o.Symbol, o.Side, o.Type, o.Price, o.Quantity, "open",
	)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	o.ID = int(id)
	o.Status = "open"

	if _, exists := me.orderBook[o.Symbol]; !exists {
		me.orderBook[o.Symbol] = NewOrderBook(o.Symbol)
	}
	me.orderBook[o.Symbol].AddOrder(o)

	me.MatchOrders(o.Symbol)
	return nil
}

func (me *MatchingEngine) CancelOrder(orderID int) error {
	_, err := me.db.Exec("UPDATE orders SET status='cancelled' WHERE id=?", orderID)
	return err
}

func (me *MatchingEngine) MatchOrders(symbol string) {
	ob, exists := me.orderBook[symbol]
	if !exists {
		return
	}

	for {
		bestBid := ob.Bids.Top()
		bestAsk := ob.Asks.Top()

		if bestBid == nil || bestAsk == nil {
			break
		}

		if bestBid.Price < bestAsk.Price {
			break
		}

		qty := min(bestBid.Quantity, bestAsk.Quantity)
		price := bestAsk.Price

		_, err := me.db.Exec(
			"INSERT INTO trades (buy_order_id, sell_order_id, symbol, price, quantity) VALUES (?, ?, ?, ?, ?)",
			bestBid.ID, bestAsk.ID, symbol, price, qty,
		)
		if err != nil {
			log.Println("Failed to insert trade:", err)
			break
		}

		bestBid.Quantity -= qty
		bestAsk.Quantity -= qty

		if bestBid.Quantity == 0 {
			me.db.Exec("UPDATE orders SET status='filled' WHERE id=?", bestBid.ID)
			ob.Bids.Pop()
		} else {
			me.db.Exec("UPDATE orders SET status='partially_filled' WHERE id=?", bestBid.ID)
		}

		if bestAsk.Quantity == 0 {
			me.db.Exec("UPDATE orders SET status='filled' WHERE id=?", bestAsk.ID)
			ob.Asks.Pop()
		} else {
			me.db.Exec("UPDATE orders SET status='partially_filled' WHERE id=?", bestAsk.ID)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (me *MatchingEngine) DB() *sql.DB {
	return me.db
}

func (me *MatchingEngine) GetOrderBook(symbol string) (*OrderBook, bool) {
	ob, ok := me.orderBook[symbol]
	return ob, ok
}
