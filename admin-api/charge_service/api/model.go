package handlers

const filePath = "../../smart-api/data/charges.json"

type Transaction struct {
	ID       uint    `json:"id"`
	UserID   string  `json:"user_id"`
	Amount   float64 `json:"amount"`
	Date     string  `json:"date"`
	Category string  `json:"category"`
	Paid     bool    `json:"paid"`
}