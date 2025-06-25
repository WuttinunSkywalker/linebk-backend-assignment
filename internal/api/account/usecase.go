package account

import "golang.org/x/sync/errgroup"

type AccountUsecase interface {
	GetAccountsByUserID(userID string, limit, offset int) ([]*AccountResponse, int, error)
}

type accountUsecase struct {
	repo AccountRepository
}

func NewAccountUsecase(repo AccountRepository) AccountUsecase {
	return &accountUsecase{
		repo: repo,
	}
}

func (u *accountUsecase) GetAccountsByUserID(userID string, limit, offset int) ([]*AccountResponse, int, error) {
	var g errgroup.Group
	var accounts []*Account
	var total int

	g.Go(func() error {
		var err error
		accounts, err = u.repo.GetAccountsByUserID(userID, limit, offset)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		total, err = u.repo.CountAccountsByUserID(userID)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, 0, err
	}

	accountResponses := make([]*AccountResponse, 0, len(accounts))
	for _, a := range accounts {
		accountResponses = append(accountResponses, NewAccountResponse(a))
	}

	return accountResponses, total, nil
}
