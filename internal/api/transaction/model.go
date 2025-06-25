package transaction

type Transaction struct {
	TransactionID string `db:"transaction_id"`
	UserID        string `db:"user_id"`
	Name          string `db:"name"`
	Image         string `db:"image"`
	IsBank        bool   `db:"is_bank"`
	CreatedAt     string `db:"created_at"`
}
