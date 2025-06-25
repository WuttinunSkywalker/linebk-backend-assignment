package transaction

import (
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"golang.org/x/sync/errgroup"
)

type TransactionUsecase interface {
	GetTransactionsByUserID(userID string, limit, offset int) ([]*TransactionResponse, int, error)
}

type transactionUsecase struct {
	transactionRepository TransactionRepository
}

func NewTransactionUsecase(transactionRepository TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepository,
	}
}

func (u *transactionUsecase) GetTransactionsByUserID(userID string, limit, offset int) ([]*TransactionResponse, int, error) {
	var g errgroup.Group
	var transaction []*Transaction
	var total int

	g.Go(func() error {
		var err error
		transaction, err = u.transactionRepository.GetTransactionsByUserID(userID, limit, offset)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		total, err = u.transactionRepository.CountTransactionsByUserID(userID)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, 0, errs.Internal(err)
	}

	transactionResponses := make([]*TransactionResponse, 0, len(transaction))
	for _, t := range transaction {
		transactionResponses = append(transactionResponses, NewTransactionResponse(t))
	}

	return transactionResponses, total, nil
}
