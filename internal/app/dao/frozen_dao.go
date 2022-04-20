package dao

import (
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"gorm.io/gorm"
)

type IAccountFrozenDao interface {
	Create(frozen *model.SpotAccountFrozen) error
	UpdateUnfreeze(req *model.UnfreezeReq) (bool, error)
	Get(orderId int64, bizType string) (*model.SpotAccountFrozen, error)
}

type AccountFrozenDao struct {
	db *gorm.DB
}

func NewAccountFrozenDao() *AccountFrozenDao {
	return &AccountFrozenDao{db: datasource.GetDB()}
}

func (d *AccountFrozenDao) Create(frozen *model.SpotAccountFrozen) error {
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
