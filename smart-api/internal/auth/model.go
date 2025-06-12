package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       string
	Username string
	Password string
}

type Charge struct {
	ID       int    `json:"id"`
	UserId   string `json:"user_id"`
	Amount   int    `json:"amount"`
	Date     string `json:"date"`
	Category string `json:"category"`
	Paid     bool   `json:"paid"`
}
