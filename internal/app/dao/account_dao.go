package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"account-module/pkg/idgen"
	"account-module/pkg/utils"
	errs "errors"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ISpotAccountDao interface {
	Save(userId int64, currency string) (bool, error)
	CreateAccountList(userIds []int64, currencies []string) error
	GetExistsAccounts(userIds []int64, currency string) ([]int64, error)
	GetAccount(userId int64, currency string) (*model.SpotAccount, error)
	GetAccountsByUserId(userId int64) ([]model.SpotAccount, error)
	HasBalance(userId int64, currency string, amount decimal.Decimal) (bool, error)
	GetLockedAccount(userId int64, currency string) (*model.SpotAccount, error)
	FreezeByUser(req *model.FreezeReq) (bool, error)
	UnfreezeByUser(req *model.UnfreezeReq) (bool, error)
	DepositByUser(req *model.DepositReq) (bool, error)
	WithdrawByUser(req *model.WithdrawReq) (bool, error)
}

type SpotAccountDao struct {
	db          *gorm.DB
	frozenDao   IAccountFrozenDao
	unfreezeDao IAccountUnFreezeDao
	logDao      IAccountLogDao
}

func NewSpotAccountDao() *SpotAccountDao {
	return &SpotAccountDao{db: datasource.GetDB()}
}

func (d *SpotAccountDao) Save(userId int64, currency string) (bool, error) {
	err := d.db.Create(&model.SpotAccount{
		AccountId: idgen.Get().GetUID(),
		UserId:    userId,
		Currency:  currency,
		Balance:   decimal.Zero,
		Frozen:    decimal.Zero}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func (d *SpotAccountDao) CreateAccountList(userIds []int64, currencies []string) error {
	uidLen := len(userIds)
	currencyLen := len(currencies)
	if uidLen <= 0 || currencyLen <= 0 {
		return errors.New("user ids or  currency list are empty")
	}

	return d.db.Transaction(func(db *gorm.DB) error {
		for _, v := range userIds {
			accountList := d.getAccountList(v, currencies)
			tx := db.CreateInBatches(accountList, currencyLen)
			if tx.Error != nil {
				return tx.Error
			}
		}

		return nil
	})
}

func (d *SpotAccountDao) GetExistsAccounts(userIds []int64, currency string) ([]int64, error) {
	var accounts []model.SpotAccount
	var resUserIds []int64
	query := "user_id IN ? AND currency = ?"
	err := d.db.Where(query, utils.Int642String(userIds), currency).Find(&accounts).Error

	for k, v := range accounts {
		resUserIds[k] = v.UserId
	}

	return resUserIds, err
}

func (d *SpotAccountDao) GetAccount(userId int64, currency string) (*model.SpotAccount, error) {
	var account model.SpotAccount

	query := "user_id = ? AND currency = ?"
	err := d.db.Where(query, userId, currency).Find(&account).Error

	return &account, err
}

func (d *SpotAccountDao) GetAccountsByUserId(userId int64) ([]model.SpotAccount, error) {
	var accounts []model.SpotAccount
	query := "user_id = ?"
	err := d.db.Where(query, userId).Find(&accounts).Error
	return accounts, err
}

func (d *SpotAccountDao) HasBalance(userId int64, currency string, amount decimal.Decimal) (bool, error) {
	account, err := d.GetAccount(userId, currency)

	if errs.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	balance := account.Balance

	return balance.GreaterThanOrEqual(amount), nil
}

func (d *SpotAccountDao) GetLockedAccount(userId int64, currency string) (*model.SpotAccount, error) {
	var account model.SpotAccount

	query := "user_id = ? AND currency = ?"
	err := d.db.Set("gorm:query_option", "FOR UPDATE").Find(&account, query, userId, currency).Error

	return &account, err
}

func (d *SpotAccountDao) FreezeByUser(req *model.FreezeReq) (bool, error) {
	query := "user_id = ? AND currency = ? AND balance >= ?"
	result := d.db.Model(&model.SpotAccount{}).
		Where(query, req.UserId, req.Currency, req.Amount).
		Updates(map[string]interface{}{
			"frozen":  gorm.Expr("frozen + ?", req.Amount),
			"balance": gorm.Expr("balance - ?", req.Amount),
		})

	return result.RowsAffected >= 1, result.Error
}

func (d *SpotAccountDao) UnfreezeByUser(req *model.UnfreezeReq) (bool, error) {
	query := "user_id = ? AND currency = ? AND frozen >= ?"
	result := d.db.Model(&model.SpotAccount{}).
		Where(query, req.UserId, req.Currency, req.Amount).
		Updates(map[string]interface{}{
			"frozen":  gorm.Expr("frozen - ?", req.Amount),
			"balance": gorm.Expr("balance + ?", req.Amount),
		})
	return result.RowsAffected >= 1, result.Error
}

func (d *SpotAccountDao) DepositByUser(req *model.DepositReq) (bool, error) {
	query := "user_id = ? AND currency = ?"
	result := d.db.Model(&model.SpotAccount{}).
		Where(query, req.UserId, req.Currency, req.Amount).
		Updates(map[string]interface{}{
			"balance": gorm.Expr("balance + ?", req.Amount),
		})

	return result.RowsAffected >= 1, result.Error
}

func (d *SpotAccountDao) WithdrawByUser(req *model.WithdrawReq) (bool, error) {
	query := "user_id = ? AND currency = ?"
	result := d.db.Model(&model.SpotAccount{}).
		Where(query, req.UserId, req.Currency, req.Amount).
		Updates(map[string]interface{}{
			"balance": gorm.Expr("balance - ?", req.Amount),
		})

	return result.RowsAffected >= 1, result.Error
}

func (d *SpotAccountDao) getAccountList(userId int64, currencies []string) []model.SpotAccount {
	var accountList []model.SpotAccount

	for _, v := range currencies {
		accountList = append(accountList, model.SpotAccount{
			AccountId: idgen.Get().GetUID(),
			UserId:    userId,
			Currency:  v,
			Balance:   decimal.Zero,
			Frozen:    decimal.Zero})
	}

	return accountList
}
