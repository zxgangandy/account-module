package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"gorm.io/gorm"
)

type IAccountUnFreezeDao interface {
	Create(frozen *model.SpotAccountUnfrozen) error
}

type AccountUnfreezeDao struct {
	db *gorm.DB
}

func NewAccountUnfreezeDao() *AccountUnfreezeDao {
	return &AccountUnfreezeDao{db: datasource.GetDB()}
}

func (s *AccountUnfreezeDao) Create(unfrozen *model.SpotAccountUnfrozen) error {
	err := s.db.Create(unfrozen).Error

	if err != nil {
		return err
	}
	return nil
}
