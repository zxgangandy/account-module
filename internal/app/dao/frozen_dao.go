package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"gorm.io/gorm"
)

type IAccountFrozenDao interface {
	Create(frozen *model.SpotAccountFrozen) (bool, error)
}

type AccountFrozenDao struct {
	db *gorm.DB
}

func NewAccountFrozenDao() *AccountFrozenDao {
	return &AccountFrozenDao{db: datasource.GetDB()}
}

func (s *AccountFrozenDao) Create(frozen *model.SpotAccountFrozen) (bool, error) {
	err := s.db.Create(frozen).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
