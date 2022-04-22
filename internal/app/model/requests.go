package model

import "github.com/shopspring/decimal"

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

type GetAccountReq struct {
	UserId   int64  `json:"userId"`
	Currency string `json:"currency"`
}

type GetAccountsReq struct {
	UserId int64 `json:"userId"`
}

type HasBalanceReq struct {
	UserId   int64           `json:"userId"`
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"currency"`
}

type FreezeReq struct {
	OrderId  int64           `json:"orderId"`
	UserId   int64           `json:"userId"`
	Currency string          `json:"currency"`
	BizType  string          `json:"bizType"`
	Amount   decimal.Decimal `json:"currency"`
}

type UnfreezeReq struct {
	OrderId  int64           `json:"orderId"`
	UserId   int64           `json:"userId"`
	Currency string          `json:"currency"`
	BizType  string          `json:"bizType"`
	Amount   decimal.Decimal `json:"currency"`
}

type DepositReq struct {
	OrderId  int64           `json:"orderId"`
	UserId   int64           `json:"userId"`
	Currency string          `json:"currency"`
	BizType  string          `json:"bizType"`
	Amount   decimal.Decimal `json:"currency"`
}

type WithdrawReq struct {
	OrderId  int64           `json:"orderId"`
	UserId   int64           `json:"userId"`
	Currency string          `json:"currency"`
	BizType  string          `json:"bizType"`
	Amount   decimal.Decimal `json:"currency"`
}

type TransferReq struct {
	OrderId    int64           `json:"orderId"`
	FromUserId int64           `json:"fromUserId"`
	ToUserId   int64           `json:"toUserId"`
	Currency   string          `json:"currency"`
	BizType    string          `json:"bizType"`
	Amount     decimal.Decimal `json:"currency"`
}
