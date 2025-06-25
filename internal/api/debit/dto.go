package debit

import (
	"fmt"

	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/logger"
)

type DebitCardResponse struct {
	CardID    string `json:"card_id" example:"card_e2c27e2b-0a8b-4791-8cf5-2186626436e5"`
	UserID    string `json:"user_id" example:"25799236-5dbc-411a-8641-b8336e088b9b"`
	Name      string `json:"name" example:"My Salary"`
	CreatedAt string `json:"created_at" example:"2024-07-04T12:40:34Z"`

	Status DebitCardStatusResponse `json:"debit_card_status"`
	Detail DebitCardDetailResponse `json:"debit_card_detail"`
	Design DebitCardDesignResponse `json:"debit_card_design"`
}

type DebitCardStatusResponse struct {
	Status string `json:"status" example:"Active"`
}

type DebitCardDetailResponse struct {
	Issuer string `json:"issuer" example:"TestLab"`
	Number string `json:"number" example:"1234 56** **** 9323"`
}

type DebitCardDesignResponse struct {
	Color       string `json:"color" example:"#00a1e2"`
	BorderColor string `json:"border_color" example:"#ffffff"`
}

func maskCardNumber(fullNumber string) string {
	logger.Info(fullNumber)
	if len(fullNumber) == 19 {
		return fmt.Sprintf("%s %s** **** %s",
			fullNumber[:4],
			fullNumber[4:6],
			fullNumber[12:16])
	} else if len(fullNumber) > 4 {
		lastFour := fullNumber[len(fullNumber)-4:]
		return fmt.Sprintf("************%s", lastFour)
	}

	return "****"
}

func NewDebitCardResponse(card *DebitCard) *DebitCardResponse {
	return &DebitCardResponse{
		CardID:    card.CardID,
		UserID:    card.UserID,
		Name:      card.Name,
		CreatedAt: card.CreatedAt,
		Status: DebitCardStatusResponse{
			Status: card.DebitCardStatus.Status,
		},
		Detail: DebitCardDetailResponse{
			Issuer: card.DebitCardDetail.Issuer,
			Number: maskCardNumber(card.DebitCardDetail.Number),
		},
		Design: DebitCardDesignResponse{
			Color:       card.DebitCardDesign.Color,
			BorderColor: card.DebitCardDesign.BorderColor,
		},
	}

}
