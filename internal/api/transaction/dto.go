package transaction

type TransactionResponse struct {
	TransactionID string `json:"transaction_id" example:"txn_a33627c4-86dd-4714-8386-29b9f9f07c00"`
	UserID        string `json:"user_id" example:"25799236-5dbc-411a-8641-b8336e088b9b"`
	Name          string `json:"name" example:"Transaction_135017"`
	Image         string `json:"image" example:"https://dummyimage.com/54x54/999/fff"`
	IsBank        bool   `json:"is_bank" example:"true"`
	CreatedAt     string `json:"created_at" example:"2024-11-23T17:18:21Z"`
}

func NewTransactionResponse(transaction *Transaction) *TransactionResponse {
	return &TransactionResponse{
		TransactionID: transaction.TransactionID,
		UserID:        transaction.UserID,
		Name:          transaction.Name,
		Image:         transaction.Image,
		IsBank:        transaction.IsBank,
		CreatedAt:     transaction.CreatedAt,
	}
}
