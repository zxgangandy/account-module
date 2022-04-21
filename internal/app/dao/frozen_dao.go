package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"gorm.io/gorm"
)

type IAccountFrozenDao interface {
	Save(frozen *model.SpotAccountFrozen) error
	UpdateUnfreeze(req *model.UnfreezeReq) (bool, error)
	Get(orderId int64, bizType string) (*model.SpotAccountFrozen, error)
	CreateFreezeOrder(account *model.SpotAccount, req *model.FreezeReq) *model.SpotAccountFrozen
}

type AccountFrozenDao struct {
	db *gorm.DB
}

func NewAccountFrozenDao() *AccountFrozenDao {
	return &AccountFrozenDao{db: datasource.GetDB()}
}

func (d *AccountFrozenDao) Save(frozen *model.SpotAccountFrozen) error {
	return d.db.Create(frozen).Error
}

func (d *AccountFrozenDao) UpdateUnfreeze(req *model.UnfreezeReq) (bool, error) {
	query := "order_id = ? AND biz_type = ? AND user_id = ? AND left_frozen >= ?"
	result := d.db.Model(&model.SpotAccountFrozen{}).
		Where(query, req.OrderId, req.BizType, req.UserId, req.Amount).
		Updates(map[string]interface{}{
			"left_frozen": gorm.Expr("left_frozen - ?", req.Amount),
		})

	return result.RowsAffected >= 1, result.Error
}

func (d *AccountFrozenDao) Get(orderId int64, bizType string) (*model.SpotAccountFrozen, error) {
	var frozen model.SpotAccountFrozen

	query := "order_id = ? AND biz_type = ?"
	err := d.db.Where(query, orderId, bizType).Find(&frozen).Error

	return &frozen, err
}

func (d *AccountFrozenDao) CreateFreezeOrder(account *model.SpotAccount, req *model.FreezeReq) *model.SpotAccountFrozen {
	return &model.SpotAccountFrozen{
		UserId:       account.UserId,
		Currency:     account.Currency,
		AccountId:    account.AccountId,
		OrderId:      req.OrderId,
		BizType:      req.BizType,
		OriginFrozen: req.Amount,
		LeftFrozen:   req.Amount,
	}
}
