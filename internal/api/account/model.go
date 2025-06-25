package account

type Account struct {
	AccountID     string `db:"account_id"`
	UserID        string `db:"user_id"`
	Name          string `db:"name"`
	Type          string `db:"type"`
	Currency      string `db:"currency"`
	AccountNumber string `db:"account_number"`
	Issuer        string `db:"issuer"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`

	AccountBalance AccountBalance `db:"account_balance"`
	AccountDetail  AccountDetail  `db:"account_detail"`
	AccountFlags   []AccountFlag  `db:"account_flag"`
}

type AccountBalance struct {
	AccountID string  `db:"account_id"`
	Amount    float64 `db:"amount"`
	UpdatedAt string  `db:"updated_at"`
}

type AccountDetail struct {
	AccountID     string `db:"account_id"`
	Color         string `db:"color"`
	IsMainAccount bool   `db:"is_main_account"`
	Progress      int    `db:"progress"`
	UpdatedAt     string `db:"updated_at"`
}

type AccountFlag struct {
	FlagID    int    `db:"flag_id"`
	AccountID string `db:"account_id"`
	FlagType  string `db:"flag_type"`
	FlagValue string `db:"flag_value"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
