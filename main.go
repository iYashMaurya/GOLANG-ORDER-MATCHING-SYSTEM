package main

import (
	"Order-Matching-System/config"
	"Order-Matching-System/database"
	"Order-Matching-System/handler"
	"Order-Matching-System/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	if err := database.CreateTables(db); err != nil {
		log.Fatal("Failed to create tables:", err)
	}

	engine := service.NewMatchingEngine(db)
	if err := engine.LoadOpenOrders(); err != nil {
		log.Fatal("Failed to load open orders:", err)
	}

	r := mux.NewRouter()
	handler.RegisterRoutes(r, engine)

	log.Printf("Server running on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
