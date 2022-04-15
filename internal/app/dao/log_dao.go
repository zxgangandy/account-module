package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"gorm.io/gorm"
)

type IAccountLogDao interface {
	Create(frozen *model.SpotAccountLog) error
}

type AccountLogDao struct {
	db *gorm.DB
}

func NewAccountLogDao() *AccountLogDao {
	return &AccountLogDao{db: datasource.GetDB()}
}

func (s *AccountLogDao) Create(log *model.SpotAccountLog) error {
	err := s.db.Create(log).Error
	if err != nil {
		return err
	}

	return nil
}
