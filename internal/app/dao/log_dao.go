package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"gorm.io/gorm"
)

type IAccountLogDao interface {
	Save(frozen *model.SpotAccountLog) error
	CreateFreezeLog(account *model.SpotAccount, req *model.FreezeReq) *model.SpotAccountLog
	CreateUnfreezeLog(account *model.SpotAccount, req *model.UnfreezeReq) *model.SpotAccountLog
	CreateDepositLog(account *model.SpotAccount, req *model.DepositReq) *model.SpotAccountLog
	CreateWithdrawLog(account *model.SpotAccount, req *model.WithdrawReq) *model.SpotAccountLog
}

type AccountLogDao struct {
	db *gorm.DB
}

func NewAccountLogDao() *AccountLogDao {
	return &AccountLogDao{db: datasource.GetDB()}
}

func (d *AccountLogDao) Save(log *model.SpotAccountLog) error {
	return d.db.Create(log).Error
}

func (d *AccountLogDao) CreateFreezeLog(account *model.SpotAccount, req *model.FreezeReq) *model.SpotAccountLog {
	return &model.SpotAccountLog{
		FromUserId:    account.UserId,
		ToUserId:      account.UserId,
		Currency:      account.Currency,
		FromAccountId: account.AccountId,
		ToAccountId:   account.AccountId,
		OrderId:       req.OrderId,
		BizType:       req.BizType,
		BeforeBalance: account.Balance,
		Balance:       account.Balance.Sub(req.Amount),
		BeforeFrozen:  account.Frozen,
		Frozen:        account.Frozen.Add(req.Amount),
		Amount:        req.Amount,
	}
}

func (d *AccountLogDao) CreateUnfreezeLog(account *model.SpotAccount, req *model.UnfreezeReq) *model.SpotAccountLog {
	return &model.SpotAccountLog{
		FromUserId:    account.UserId,
		ToUserId:      account.UserId,
		Currency:      account.Currency,
		FromAccountId: account.AccountId,
		ToAccountId:   account.AccountId,
		OrderId:       req.OrderId,
		BizType:       req.BizType,
		BeforeBalance: account.Balance,
		Balance:       account.Balance.Add(req.Amount),
		BeforeFrozen:  account.Frozen,
		Frozen:        account.Frozen.Sub(req.Amount),
		Amount:        req.Amount,
	}
}

func (d *AccountLogDao) CreateDepositLog(account *model.SpotAccount, req *model.DepositReq) *model.SpotAccountLog {
	return &model.SpotAccountLog{
		FromUserId:    account.UserId,
		ToUserId:      account.UserId,
		Currency:      account.Currency,
		FromAccountId: account.AccountId,
		ToAccountId:   account.AccountId,
		OrderId:       req.OrderId,
		BizType:       req.BizType,
		BeforeBalance: account.Balance,
		Balance:       account.Balance.Add(req.Amount),
		BeforeFrozen:  account.Frozen,
		Frozen:        account.Frozen,
		Amount:        req.Amount,
	}
}

func (d *AccountLogDao) CreateWithdrawLog(account *model.SpotAccount, req *model.WithdrawReq) *model.SpotAccountLog {
	return &model.SpotAccountLog{
		FromUserId:    account.UserId,
		ToUserId:      account.UserId,
		Currency:      account.Currency,
		FromAccountId: account.AccountId,
		ToAccountId:   account.AccountId,
		OrderId:       req.OrderId,
		BizType:       req.BizType,
		BeforeBalance: account.Balance,
		Balance:       account.Balance.Sub(req.Amount),
		BeforeFrozen:  account.Frozen,
		Frozen:        account.Frozen,
		Amount:        req.Amount,
	}
}
