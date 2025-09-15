package models

type Trade struct {
	ID          int
	BuyOrderID  int
	SellOrderID int
	Symbol      string
	Price       float64
	Quantity    int
	CreatedAt   string
}
