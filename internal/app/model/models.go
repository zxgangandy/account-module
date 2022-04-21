package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type SpotAccount struct {
	Id        uint64          `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	AccountId int64           `gorm:"column:account_id;unique_index:uniq_account_id"`
	UserId    int64           `gorm:"column:user_id;unique_index:uniq_user_currency"`
	Currency  string          `gorm:"column:currency;unique_index:uniq_user_currency"`
	Balance   decimal.Decimal `gorm:"column:balance" sql:"type:decimal(32,16);"`
	Frozen    decimal.Decimal `gorm:"column:frozen" sql:"type:decimal(32,16);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SpotAccountFrozen struct {
	Id           uint64          `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserId       int64           `gorm:"column:user_id"`
	Currency     string          `gorm:"column:currency"`
	AccountId    int64           `gorm:"column:account_id"`
	OrderId      int64           `gorm:"column:account_id;unique_index:uniq_biz_order"`
	BizType      string          `gorm:"column:biz_type;unique_index:uniq_biz_order"`
	OriginFrozen decimal.Decimal `gorm:"column:origin_frozen" sql:"type:decimal(32,16);"`
	LeftFrozen   decimal.Decimal `gorm:"column:left_frozen" sql:"type:decimal(32,16);"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SpotAccountUnfrozen struct {
	Id           uint64          `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	BizId        int64           `gorm:"column:biz_id"`
	UserId       int64           `gorm:"column:user_id"`
	Currency     string          `gorm:"column:currency"`
	AccountId    int64           `gorm:"column:account_id"`
	OrderId      int64           `gorm:"column:account_id;unique_index:uniq_biz_order"`
	BizType      string          `gorm:"column:biz_type;unique_index:uniq_biz_order"`
	OriginFrozen decimal.Decimal `gorm:"column:origin_frozen" sql:"type:decimal(32,16);"`
	LeftFrozen   decimal.Decimal `gorm:"column:left_frozen" sql:"type:decimal(32,16);"`
	Unfrozen     decimal.Decimal `gorm:"column:frozen" sql:"type:decimal(32,16);"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SpotAccountTrade struct {
	Id            uint64          `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserId        int64           `gorm:"column:user_id"`
	Currency      string          `gorm:"column:currency"`
	AccountId     int64           `gorm:"column:account_id"`
	OrderId       int64           `gorm:"column:account_id"`
	BizType       string          `gorm:"column:biz_type"`
	TradeType     string          `gorm:"column:trade_type"`
	BeforeBalance decimal.Decimal `gorm:"column:before_balance" sql:"type:decimal(32,16);"`
	AfterBalance  decimal.Decimal `gorm:"column:after_balance" sql:"type:decimal(32,16);"`
	Amount        decimal.Decimal `gorm:"column:amount" sql:"type:decimal(32,16);"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type SpotAccountLog struct {
	Id            uint64          `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	FromUserId    int64           `gorm:"column:from_user_id"`
	ToUserId      int64           `gorm:"column:to_user_id"`
	Currency      string          `gorm:"column:currency"`
	FromAccountId int64           `gorm:"column:from_account_id"`
	ToAccountId   int64           `gorm:"column:to_account_id"`
	OrderId       int64           `gorm:"column:account_id"`
	BizType       string          `gorm:"column:biz_type"`
	BeforeBalance decimal.Decimal `gorm:"column:before_balance" sql:"type:decimal(32,16);"`
	Balance       decimal.Decimal `gorm:"column:balance" sql:"type:decimal(32,16);"`
	BeforeFrozen  decimal.Decimal `gorm:"column:before_frozen" sql:"type:decimal(32,16);"`
	Frozen        decimal.Decimal `gorm:"column:frozen" sql:"type:decimal(32,16);"`
	Amount        decimal.Decimal `gorm:"column:amount" sql:"type:decimal(32,16);"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
