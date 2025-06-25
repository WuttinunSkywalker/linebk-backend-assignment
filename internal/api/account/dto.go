package account

type AccountResponse struct {
	AccountID     string `json:"account_id" example:"acc_4d2fbe84-89ee-4d6f-8088-46b0121000ed"`
	UserID        string `json:"user_id" example:"25799236-5dbc-411a-8641-b8336e088b9b"`
	Name          string `json:"name" example:"Saving Account"`
	Type          string `json:"type" example:"saving-account"`
	Currency      string `json:"currency" example:"THB"`
	AccountNumber string `json:"account_number" example:"090-5-70425-2"`
	Issuer        string `json:"issuer" example:"TestLab"`
	CreatedAt     string `json:"created_at" example:"2025-03-24T13:12:39Z"`

	AccountBalance AccountBalanceResponse `json:"account_balance"`
	AccountDetail  AccountDetailResponse  `json:"account_detail"`
	AccountFlags   []AccountFlagResponse  `json:"account_flag"`
}

type AccountBalanceResponse struct {
	Amount float64 `json:"amount" example:"181471.73"`
}

type AccountDetailResponse struct {
	Color         string `json:"color" example:"#00a1e2"`
	IsMainAccount bool   `json:"is_main_account" example:"true"`
	Progress      int    `json:"progress" example:"0"`
}

type AccountFlagResponse struct {
	FlagID    int    `json:"flag_id" example:"1"`
	FlagType  string `json:"flag_type" example:"system"`
	FlagValue string `json:"flag_value" example:"Disbursement"`
}

func NewAccountResponse(acct *Account) *AccountResponse {
	flagDTOs := make([]AccountFlagResponse, 0, len(acct.AccountFlags))
	for _, f := range acct.AccountFlags {
		flagDTOs = append(flagDTOs, AccountFlagResponse{
			FlagID:    f.FlagID,
			FlagType:  f.FlagType,
			FlagValue: f.FlagValue,
		})
	}

	return &AccountResponse{
		AccountID:     acct.AccountID,
		UserID:        acct.UserID,
		Name:          acct.Name,
		Type:          acct.Type,
		Currency:      acct.Currency,
		AccountNumber: acct.AccountNumber,
		Issuer:        acct.Issuer,
		CreatedAt:     acct.CreatedAt,
		AccountBalance: AccountBalanceResponse{
			Amount: acct.AccountBalance.Amount,
		},
		AccountDetail: AccountDetailResponse{
			Color:         acct.AccountDetail.Color,
			IsMainAccount: acct.AccountDetail.IsMainAccount,
			Progress:      acct.AccountDetail.Progress,
		},
		AccountFlags: flagDTOs,
	}
}
