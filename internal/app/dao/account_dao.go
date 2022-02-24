package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ISpotAccountDao interface {
	Create(userId uint64, currency string) error
	CreateAccountList(userIds []uint64, currencies []string) error
}

type SpotAccountDao struct {
	db *gorm.DB
}

func NewSpotAccountDao() *SpotAccountDao {
	return &SpotAccountDao{db: datasource.GetDB()}
}

func (s *SpotAccountDao) Create(userId uint64, currency string) error {
	err := s.db.Create(&model.SpotAccount{
		AccountId: 1,
		UserId:    userId,
		Currency:  currency,
		Balance:   decimal.Zero,
		Frozen:    decimal.Zero}).Error
	return err
}

func (s *SpotAccountDao) CreateAccountList(userIds []uint64, currencies []string) error {
	uidLen := len(userIds)
	currencyLen := len(currencies)
	if uidLen <= 0 || currencyLen <= 0 {
		return errors.New("user ids or  currency list are empty")
	}

	s.db.Transaction(func(db *gorm.DB) error {
		for _, v := range userIds {
			accountList := s.getAccountList(v, currencies)
			db.CreateInBatches(accountList, currencyLen)
		}

		return nil
	})

	return nil
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
