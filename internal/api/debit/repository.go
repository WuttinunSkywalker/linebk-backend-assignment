package debit

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DebitRepository interface {
	GetDebitCardsByUserID(userID string, limit, offset int) ([]*DebitCard, error)
	CountDebitCardsByUserID(userID string) (int, error)
}

type debitRepository struct {
	db *sqlx.DB
}

func NewDebitRepository(db *sqlx.DB) DebitRepository {
	return &debitRepository{
		db: db,
	}
}

func (r *debitRepository) GetDebitCardsByUserID(userID string, limit, offset int) ([]*DebitCard, error) {
	const query = `
    WITH user_cards AS (
        SELECT * FROM debit_cards
        WHERE user_id = ?
        ORDER BY created_at DESC
        LIMIT ?
        OFFSET ?
    )
    SELECT
        -- main 
        uc.card_id,
        uc.user_id,
        uc.name,
        uc.created_at,
        uc.updated_at,

        -- status
        dcs.card_id     AS "debit_card_status.card_id",
        dcs.status      AS "debit_card_status.status",
        dcs.updated_at  AS "debit_card_status.updated_at",

        -- detail
        dcd.card_id     AS "debit_card_details.card_id",
        dcd.issuer      AS "debit_card_details.issuer",
        dcd.number      AS "debit_card_details.number",
        dcd.updated_at  AS "debit_card_details.updated_at",

        -- design
        dcdn.card_id       AS "debit_card_design.card_id",
        dcdn.color         AS "debit_card_design.color",
        dcdn.border_color  AS "debit_card_design.border_color",
        dcdn.updated_at    AS "debit_card_design.updated_at"
    FROM
        user_cards uc
    JOIN
        debit_card_status dcs ON uc.card_id = dcs.card_id
    JOIN
        debit_card_details dcd ON uc.card_id = dcd.card_id
    JOIN
        debit_card_design dcdn ON uc.card_id = dcdn.card_id
    ORDER BY
        uc.created_at DESC; 
    `

	var debitCards []*DebitCard

	err := r.db.Select(&debitCards, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying for user's debit cards: %w", err)
	}

	return debitCards, nil
}

func (r *debitRepository) CountDebitCardsByUserID(userID string) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM debit_cards WHERE user_id = ?"
	err := r.db.Get(&count, query, userID)
	return count, err
}
