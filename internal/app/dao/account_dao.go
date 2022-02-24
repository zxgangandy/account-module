package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ISpotAccountDao interface {
	Create(userId uint64, currency string) (bool, error)
	CreateAccountList(userIds []uint64, currencies []string) error
}

type SpotAccountDao struct {
	db *gorm.DB
}

func NewSpotAccountDao() *SpotAccountDao {
	return &SpotAccountDao{db: datasource.GetDB()}
}

func (s *SpotAccountDao) Create(userId uint64, currency string) (bool, error) {
	err := s.db.Create(&model.SpotAccount{
		AccountId: 1,
		UserId:    userId,
		Currency:  currency,
		Balance:   decimal.Zero,
		Frozen:    decimal.Zero}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *SpotAccountDao) CreateAccountList(userIds []uint64, currencies []string) error {
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

func (s *SpotAccountDao) getAccountList(userId uint64, currencies []string) []model.SpotAccount {
	var accountList []model.SpotAccount

	for _, v := range currencies {
		accountList = append(accountList, model.SpotAccount{
			AccountId: 1,
			UserId:    userId,
			Currency:  v,
			Balance:   decimal.Zero,
			Frozen:    decimal.Zero})
	}

	return accountList
}
