package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"gorm.io/gorm"
)

const (
	IN  = "in"
	OUT = "out"
)

type IAccountTradeDao interface {
	Save(frozen *model.SpotAccountTrade) error
	CreateDepositOrder(account *model.SpotAccount, req *model.DepositReq) *model.SpotAccountTrade
	CreateWithdrawOrder(account *model.SpotAccount, req *model.WithdrawReq) *model.SpotAccountTrade
}

type AccountTradeDao struct {
	db *gorm.DB
}

func NewAccountTradeDao() *AccountTradeDao {
	return &AccountTradeDao{db: datasource.GetDB()}
}

func (d *AccountTradeDao) Save(trade *model.SpotAccountTrade) error {
	return d.db.Create(trade).Error
}

func (d *AccountTradeDao) CreateDepositOrder(account *model.SpotAccount, req *model.DepositReq) *model.SpotAccountTrade {
	return &model.SpotAccountTrade{
		UserId:        req.UserId,
		Currency:      req.Currency,
		AccountId:     account.AccountId,
		OrderId:       req.OrderId,
		BizType:       req.BizType,
		TradeType:     IN,
		BeforeBalance: account.Balance,
		AfterBalance:  account.Balance.Add(req.Amount),
		Amount:        req.Amount,
	}
}

func (d *AccountTradeDao) CreateWithdrawOrder(account *model.SpotAccount, req *model.WithdrawReq) *model.SpotAccountTrade {
	return &model.SpotAccountTrade{
		UserId:        req.UserId,
		Currency:      req.Currency,
		AccountId:     account.AccountId,
		OrderId:       req.OrderId,
		BizType:       req.BizType,
		TradeType:     OUT,
		BeforeBalance: account.Balance,
		AfterBalance:  account.Balance.Sub(req.Amount),
		Amount:        req.Amount,
	}
}
