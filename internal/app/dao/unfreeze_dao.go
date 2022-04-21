package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"account-module/pkg/idgen"
	"gorm.io/gorm"
)

type IAccountUnFreezeDao interface {
	Save(frozen *model.SpotAccountUnfrozen) error
	CreateUnfreezeOrder(frozen *model.SpotAccountFrozen, req *model.UnfreezeReq) *model.SpotAccountUnfrozen
}

type AccountUnfreezeDao struct {
	db *gorm.DB
}

func NewAccountUnfreezeDao() *AccountUnfreezeDao {
	return &AccountUnfreezeDao{db: datasource.GetDB()}
}

func (d *AccountUnfreezeDao) Save(unfrozen *model.SpotAccountUnfrozen) error {
	return d.db.Create(unfrozen).Error
}

func (d *AccountUnfreezeDao) CreateUnfreezeOrder(frozen *model.SpotAccountFrozen, req *model.UnfreezeReq) *model.SpotAccountUnfrozen {
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
