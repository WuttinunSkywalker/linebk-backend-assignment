package debit

type DebitCard struct {
	CardID    string `db:"card_id"`
	UserID    string `db:"user_id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`

	DebitCardStatus DebitCardStatus `db:"debit_card_status"`
	DebitCardDetail DebitCardDetail `db:"debit_card_details"`
	DebitCardDesign DebitCardDesign `db:"debit_card_design"`
}

type DebitCardStatus struct {
	CardID    string `db:"card_id"`
	Status    string `db:"status"`
	UpdatedAt string `db:"updated_at"`
}

type DebitCardDetail struct {
	CardID    string `db:"card_id"`
	Issuer    string `db:"issuer"`
	Number    string `db:"number"`
	UpdatedAt string `db:"updated_at"`
}

type DebitCardDesign struct {
	CardID      string `db:"card_id"`
	Color       string `db:"color"`
	BorderColor string `db:"border_color"`
	UpdatedAt   string `db:"updated_at"`
}
