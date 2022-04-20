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

func (d *AccountUnfreezeDao) Create(unfrozen *model.SpotAccountUnfrozen) error {
	return d.db.Create(unfrozen).Error
}
