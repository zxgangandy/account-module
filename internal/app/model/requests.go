package model

type CreateAccountReq struct {
	UserID   uint64 `json:"userId"`
	Currency string `json:"currency"`
}

type CreateAccountListReq struct {
	UserIDList   []uint64 `json:"userIds"`
	CurrencyList []string `json:"currencies"`
}

type ExistAccountListReq struct {
	UserIDList []uint64 `json:"userIds"`
	Currency   string   `json:"currency"`
}
