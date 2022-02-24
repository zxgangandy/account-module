package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type SpotAccount struct {
	Id        uint64          `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	AccountId uint64          `gorm:"column:account_id;unique_index:uniq_account_id"`
	UserId    uint64          `gorm:"column:user_id;unique_index:uniq_user_currency"`
	Currency  string          `gorm:"column:currency;unique_index:uniq_user_currency"`
	Balance   decimal.Decimal `gorm:"column:balance" sql:"type:decimal(32,16);"`
	Frozen    decimal.Decimal `gorm:"column:frozen" sql:"type:decimal(32,16);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
