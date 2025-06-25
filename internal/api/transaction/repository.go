package transaction

import "github.com/jmoiron/sqlx"

type TransactionRepository interface {
	GetTransactionsByUserID(userID string, limit, offset int) ([]*Transaction, error)
	CountTransactionsByUserID(userID string) (int, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) GetTransactionsByUserID(userID string, limit, offset int) ([]*Transaction, error) {
	var transactions []*Transaction
	query := "SELECT transaction_id, user_id, name, image, is_bank, created_at FROM transactions WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?"
	err := r.db.Select(&transactions, query, userID, limit, offset)
	return transactions, err
}

func (r *transactionRepository) CountTransactionsByUserID(userID string) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM transactions WHERE user_id = ?"
	err := r.db.Get(&count, query, userID)
	return count, err
}
