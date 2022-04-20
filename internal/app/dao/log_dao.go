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

func (d *AccountLogDao) Create(log *model.SpotAccountLog) error {
	return d.db.Create(log).Error
}
