package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"container/list"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ISpotAccountDao interface {
	Create(userId int64, currency string) error
	CreateAccountList(userIds *list.List, currencies *list.List) error
}

type SpotAccountDao struct {
	db *gorm.DB
}

func NewSpotAccountDao() *SpotAccountDao {
	return &SpotAccountDao{db: datasource.GetDB()}
}

func (s *SpotAccountDao) Create(userId int64, currency string) error {
	err := s.db.Create(&model.SpotAccount{
		AccountId: 1,
		UserId:    userId,
		Currency:  currency,
		Balance:   decimal.Zero,
		Frozen:    decimal.Zero}).Error
	return err
}

func (s *SpotAccountDao) CreateAccountList(userIds *list.List, currencies *list.List) error {
	uidLen := userIds.Len()
	if uidLen <= 0 || currencies.Len() <= 0 {
		return errors.New("user ids or  currency list are empty")
	}

	//s.db.Transaction(func(db *gorm.DB) error {
	//	for k, v:= range userIds  {
	//		s.db.CreateInBatches()
	//	}
	//})

	//for i := userIds.Front(); i != nil; i = i.Next()  {
	//	s.db.CreateInBatches()
	//}

	return nil
}
