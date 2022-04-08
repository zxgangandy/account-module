package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"account-module/pkg/idgen"
	"account-module/pkg/utils"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ISpotAccountDao interface {
	Create(userId int64, currency string) (bool, error)
	CreateAccountList(userIds []int64, currencies []string) error
	GetExistsAccounts(userIds []int64, currency string) ([]int64, error)
	GetAccount(userId int64, currency string) (model.SpotAccount, error)
	GetAccountsByUserId(userId int64) ([]model.SpotAccount, error)
	HasBalance(userId int64, currency string, amount decimal.Decimal) (bool, error)
}

type SpotAccountDao struct {
	db *gorm.DB
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

func (s *SpotAccountDao) GetAccount(userId int64, currency string) (model.SpotAccount, error) {
	var account model.SpotAccount

	query := "user_id = ? AND currency = ?"
	err := s.db.Where(query, userId, currency).Find(&account).Error

	return account, err
}

func (s *SpotAccountDao) GetAccountsByUserId(userId int64) ([]model.SpotAccount, error) {
	var accounts []model.SpotAccount
	query := "user_id = ?"
	err := s.db.Where(query, userId).Find(&accounts).Error
	return accounts, err
}

func (s *SpotAccountDao) HasBalance(userId int64, currency string, amount decimal.Decimal) (bool, error) {
	account, err := s.GetAccount(userId, currency)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	balance := account.Balance

	return balance.GreaterThanOrEqual(amount), nil
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
