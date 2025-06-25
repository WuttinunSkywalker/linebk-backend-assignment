package account

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AccountRepository interface {
	GetAccountsByUserID(userID string, limit, offset int) ([]*Account, error)
	CountAccountsByUserID(userID string) (int, error)
}

type accountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) GetAccountsByUserID(userID string, limit, offset int) ([]*Account, error) {
	const baseQuery = `
	WITH user_accounts AS (
		SELECT *
		FROM   accounts
		WHERE  user_id = ?
		ORDER  BY created_at DESC
		LIMIT  ?
		OFFSET ?
	)
	SELECT
		-- main
		ua.account_id,
		ua.user_id,
		ua.name,
		ua.type,
		ua.currency,
		ua.account_number,
		ua.issuer,
		ua.created_at,
		ua.updated_at,

		-- balance
		ab.account_id AS "account_balance.account_id",
		ab.amount     AS "account_balance.amount",
		ab.updated_at AS "account_balance.updated_at",

		-- detail
		ad.account_id      AS "account_detail.account_id",
		ad.color           AS "account_detail.color",
		ad.is_main_account AS "account_detail.is_main_account",
		ad.progress        AS "account_detail.progress",
		ad.updated_at      AS "account_detail.updated_at"
	FROM
		user_accounts ua
	JOIN account_balances ab ON ua.account_id = ab.account_id
	JOIN account_details ad ON ua.account_id = ad.account_id
	ORDER BY is_main_account DESC, ua.created_at DESC;
	`

	var accounts []*Account
	if err := r.db.Select(&accounts, baseQuery, userID, limit, offset); err != nil {
		return nil, fmt.Errorf("query accounts: %w", err)
	}

	if len(accounts) == 0 {
		return accounts, nil
	}

	ids := make([]interface{}, 0, len(accounts))
	for _, a := range accounts {
		ids = append(ids, a.AccountID)
	}

	flagsSQL, args, err := sqlx.In(`
		SELECT *
		FROM   account_flags
		WHERE  account_id IN (?);`, ids)
	if err != nil {
		return nil, fmt.Errorf("build flags query: %w", err)
	}

	var flags []AccountFlag
	if err := r.db.Select(&flags, flagsSQL, args...); err != nil {
		return nil, fmt.Errorf("query flags: %w", err)
	}

	flagMap := make(map[string][]AccountFlag, len(ids))
	for _, f := range flags {
		flagMap[f.AccountID] = append(flagMap[f.AccountID], f)
	}
	for i := range accounts {
		accounts[i].AccountFlags = flagMap[accounts[i].AccountID]
	}

	return accounts, nil
}

func (r *accountRepository) CountAccountsByUserID(userID string) (int, error) {
	const query = `SELECT COUNT(*) FROM accounts WHERE user_id = ?`
	var count int
	err := r.db.Get(&count, query, userID)
	return count, err
}
