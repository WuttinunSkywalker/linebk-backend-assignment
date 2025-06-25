package debit

import (
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/errs"
	"golang.org/x/sync/errgroup"
)

type DebitUseCase interface {
	GetDebitCardsByUserID(userID string, limit, offset int) ([]*DebitCardResponse, int, error)
}

type debitUsecase struct {
	repo DebitRepository
}

func NewDebitUsecase(repo DebitRepository) DebitUseCase {
	return &debitUsecase{
		repo: repo,
	}
}

func (u *debitUsecase) GetDebitCardsByUserID(userID string, limit, offset int) ([]*DebitCardResponse, int, error) {
	var g errgroup.Group
	var debitCards []*DebitCard
	var total int

	g.Go(func() error {
		var err error
		debitCards, err = u.repo.GetDebitCardsByUserID(userID, limit, offset)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		total, err = u.repo.CountDebitCardsByUserID(userID)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, 0, errs.Internal(err)
	}

	debitCardResponses := make([]*DebitCardResponse, 0, len(debitCards))
	for _, d := range debitCards {
		debitCardResponses = append(debitCardResponses, NewDebitCardResponse(d))
	}

	return debitCardResponses, total, nil
}
