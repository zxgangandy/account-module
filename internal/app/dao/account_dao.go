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
	Create(userId int64, currency string) (bool, error)
	CreateAccountList(userIds []int64, currencies []string) error
	GetExistsAccounts(userIds []int64, currency string) ([]int64, error)
	GetAccount(userId int64, currency string) (*model.SpotAccount, error)
	GetAccountsByUserId(userId int64) ([]model.SpotAccount, error)
	HasBalance(userId int64, currency string, amount decimal.Decimal) (bool, error)
	GetLockedAccount(userId int64, currency string) (*model.SpotAccount, error)
	FreezeByUser(req *model.FreezeReq) (bool, error)
	CreateFreezeOrder(account *model.SpotAccount, req *model.FreezeReq) *model.SpotAccountFrozen
	CreateFreezeLog(account *model.SpotAccount, req *model.FreezeReq) *model.SpotAccountLog
	UnfreezeByUser(req *model.UnfreezeReq) (bool, error)
	CreateUnfreezeOrder(frozen *model.SpotAccountFrozen, req *model.UnfreezeReq) *model.SpotAccountUnfrozen
	CreateUnfreezeLog(account *model.SpotAccount, req *model.UnfreezeReq) *model.SpotAccountLog
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

func (s *SpotAccountDao) Create(userId int64, currency string) (bool, error) {
	err := s.db.Create(&model.SpotAccount{
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

func (s *SpotAccountDao) CreateAccountList(userIds []int64, currencies []string) error {
	uidLen := len(userIds)
	currencyLen := len(currencies)
	if uidLen <= 0 || currencyLen <= 0 {
		return errors.New("user ids or  currency list are empty")
	}

	return s.db.Transaction(func(db *gorm.DB) error {
		for _, v := range userIds {
			accountList := s.getAccountList(v, currencies)
			tx := db.CreateInBatches(accountList, currencyLen)
			if tx.Error != nil {
				return tx.Error
			}
		}

		return nil
	})
}

func (s *SpotAccountDao) GetExistsAccounts(userIds []int64, currency string) ([]int64, error) {
	var accounts []model.SpotAccount
	var resUserIds []int64
	query := "user_id IN ? AND currency = ?"
	err := s.db.Where(query, utils.Int642String(userIds), currency).Find(&accounts).Error

	for k, v := range accounts {
		resUserIds[k] = v.UserId
	}

	return resUserIds, err
}

func (s *SpotAccountDao) GetAccount(userId int64, currency string) (*model.SpotAccount, error) {
	var account model.SpotAccount

	query := "user_id = ? AND currency = ?"
	err := s.db.Where(query, userId, currency).Find(&account).Error

	return &account, err
}

func (s *SpotAccountDao) GetAccountsByUserId(userId int64) ([]model.SpotAccount, error) {
	var accounts []model.SpotAccount
	query := "user_id = ?"
	err := s.db.Where(query, userId).Find(&accounts).Error
	return accounts, err
}

func (s *SpotAccountDao) HasBalance(userId int64, currency string, amount decimal.Decimal) (bool, error) {
	account, err := s.GetAccount(userId, currency)

	if errs.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	balance := account.Balance

	return balance.GreaterThanOrEqual(amount), nil
}

func (s *SpotAccountDao) GetLockedAccount(userId int64, currency string) (*model.SpotAccount, error) {
	var account model.SpotAccount

	query := "user_id = ? AND currency = ?"
	err := s.db.Set("gorm:query_option", "FOR UPDATE").Find(&account, query, userId, currency).Error

	return &account, err
}

func (s *SpotAccountDao) FreezeByUser(req *model.FreezeReq) (bool, error) {
	query := "user_id = ? AND currency = ? AND balance >= ?"
	result := s.db.Model(&model.SpotAccount{}).
		Where(query, req.UserId, req.Currency, req.Amount).
		Updates(map[string]interface{}{
			"frozen":  gorm.Expr("frozen + ?", req.Amount),
			"balance": gorm.Expr("balance - ?", req.Amount),
		})

	return result.RowsAffected >= 1, result.Error
}

func (s *SpotAccountDao) UnfreezeByUser(req *model.UnfreezeReq) (bool, error) {
	query := "user_id = ? AND currency = ? AND frozen >= ?"
	result := s.db.Model(&model.SpotAccount{}).
		Where(query, req.UserId, req.Currency, req.Amount).
		Updates(map[string]interface{}{
			"frozen":  gorm.Expr("frozen - ?", req.Amount),
			"balance": gorm.Expr("balance + ?", req.Amount),
		})
	return result.RowsAffected >= 1, result.Error
}

func (s *SpotAccountDao) getAccountList(userId int64, currencies []string) []model.SpotAccount {
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

func (s *SpotAccountDao) CreateFreezeOrder(account *model.SpotAccount, req *model.FreezeReq) *model.SpotAccountFrozen {
	return &model.SpotAccountFrozen{
		UserId:       account.UserId,
		Currency:     account.Currency,
		AccountId:    account.AccountId,
		OrderId:      req.OrderId,
		BizType:      req.BizType,
		OriginFrozen: req.Amount,
		LeftFrozen:   req.Amount,
	}
}

func (s *SpotAccountDao) CreateFreezeLog(account *model.SpotAccount, req *model.FreezeReq) *model.SpotAccountLog {
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

func (s *SpotAccountDao) CreateUnfreezeOrder(frozen *model.SpotAccountFrozen, req *model.UnfreezeReq) *model.SpotAccountUnfrozen {
	return &model.SpotAccountUnfrozen{
		BizId:        idgen.Get().GetUID(),
		UserId:       req.UserId,
		Currency:     req.Currency,
		AccountId:    frozen.AccountId,
		OrderId:      req.OrderId,
		BizType:      req.BizType,
		OriginFrozen: frozen.OriginFrozen,
		LeftFrozen:   frozen.LeftFrozen.Sub(req.Amount),
		Unfrozen:     req.Amount,
	}
}

func (s *SpotAccountDao) CreateUnfreezeLog(account *model.SpotAccount, req *model.UnfreezeReq) *model.SpotAccountLog {
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
