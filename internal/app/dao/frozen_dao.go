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

func (s *AccountFrozenDao) Create(frozen *model.SpotAccountFrozen) error {
	err := s.db.Create(frozen).Error

	if err != nil {
		return err
	}
	return nil
}

func (s *AccountFrozenDao) UpdateUnfreeze(req *model.UnfreezeReq) (bool, error) {
	query := "order_id = ? AND biz_type = ? AND user_id = ? AND left_frozen >= ?"
	result := s.db.Model(&model.SpotAccountFrozen{}).
		Where(query, req.OrderId, req.BizType, req.UserId, req.Amount).
		Updates(map[string]interface{}{
			"left_frozen": gorm.Expr("left_frozen - ?", req.Amount),
		})

	return result.RowsAffected >= 1, result.Error
}

func (s *AccountFrozenDao) Get(orderId int64, bizType string) (*model.SpotAccountFrozen, error) {
	var frozen model.SpotAccountFrozen

	query := "order_id = ? AND biz_type = ?"
	err := s.db.Where(query, orderId, bizType).Find(&frozen).Error

	return &frozen, err
}
