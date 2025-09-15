package handler

import (
	"Order-Matching-System/models"
	"Order-Matching-System/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, engine *service.MatchingEngine) {
	log.Println("Registering routes...")
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")
	r.HandleFunc("/orders", placeOrderHandler(engine)).Methods("POST")
	r.HandleFunc("/orders/{orderId}", cancelOrderHandler(engine)).Methods("DELETE")
	r.HandleFunc("/orders/{orderId}", getOrderStatusHandler(engine)).Methods("GET")
	r.HandleFunc("/orderbook", getOrderBookHandler(engine)).Methods("GET")
	r.HandleFunc("/trades", getTradesHandler(engine)).Methods("GET")
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	writeResponse(w, http.StatusOK, map[string]string{"status": "ok"})
}

func placeOrderHandler(engine *service.MatchingEngine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var o models.Order
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			writeResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
			return
		}
		if err := engine.PlaceOrder(&o); err != nil {
			writeResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeResponse(w, http.StatusCreated, o)
	}
}

func cancelOrderHandler(engine *service.MatchingEngine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["orderId"]
		id, _ := strconv.Atoi(idStr)
		if err := engine.CancelOrder(id); err != nil {
			writeResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeResponse(w, http.StatusOK, map[string]string{"message": "order cancelled"})
	}
}

func getOrderStatusHandler(engine *service.MatchingEngine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["orderId"]
		id, _ := strconv.Atoi(idStr)

		row := engine.DB().QueryRow("SELECT id, symbol, side, type, price, quantity, status, created_at FROM orders WHERE id=?", id)
		var o models.Order
		err := row.Scan(&o.ID, &o.Symbol, &o.Side, &o.Type, &o.Price, &o.Quantity, &o.Status, &o.CreatedAt)
		if err != nil {
			writeResponse(w, http.StatusNotFound, map[string]string{"error": "order not found"})
			return
		}
		writeResponse(w, http.StatusOK, o)
	}
}

func getOrderBookHandler(engine *service.MatchingEngine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Query().Get("symbol")
		if ob, ok := engine.GetOrderBook(symbol); ok {
			writeResponse(w, http.StatusOK, ob)
			return
		}
		writeResponse(w, http.StatusNotFound, map[string]string{"error": "orderbook not found"})
	}
}

func getTradesHandler(engine *service.MatchingEngine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := engine.DB().Query("SELECT id, buy_order_id, sell_order_id, symbol, price, quantity, created_at FROM trades")
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		defer rows.Close()

		var trades []models.Trade
		for rows.Next() {
			var t models.Trade
			rows.Scan(&t.ID, &t.BuyOrderID, &t.SellOrderID, &t.Symbol, &t.Price, &t.Quantity, &t.CreatedAt)
			trades = append(trades, t)
		}
		writeResponse(w, http.StatusOK, trades)
	}
}

func writeResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
