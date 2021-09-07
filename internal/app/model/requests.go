package model

type CreateAccountReq struct {
	UserID   int64  `json:"userId"`
	Currency string `json:"currency"`
}
