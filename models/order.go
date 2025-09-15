package models

type Order struct {
	ID        int
	Symbol    string
	Side      string
	Type      string
	Price     float64
	Quantity  int
	Status    string
	CreatedAt string
}
