package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"gorm.io/gorm"
)

type IAccountTradeDao interface {
	Save(frozen *model.SpotAccountTrade) error
	CreateDepositOrder(frozen *model.SpotAccount, req *model.DepositReq) *model.SpotAccountTrade
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
		TradeType:     "in",
		BeforeBalance: account.Balance,
		AfterBalance:  account.Balance.Add(req.Amount),
		Amount:        req.Amount,
	}
}
