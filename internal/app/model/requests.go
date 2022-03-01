package model

type CreateAccountReq struct {
	UserID   int64  `json:"userId"`
	Currency string `json:"currency"`
}

type CreateAccountListReq struct {
	UserIDList   []int64  `json:"userIds"`
	CurrencyList []string `json:"currencies"`
}

type ExistAccountListReq struct {
	UserIDList []int64 `json:"userIds"`
	Currency   string  `json:"currency"`
}
