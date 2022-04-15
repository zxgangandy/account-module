package service

import (
	"account-module/internal/app/dao"
	"account-module/internal/app/model"
	"github.com/shopspring/decimal"
)

type IAccountService interface {
	CreateAccount(userId int64, currency string) (bool, error)
	CreateAccountList(userIds []int64, currencies []string) error
	GetExistsAccounts(userIds []int64, currency string) ([]int64, error)
	GetAccount(userId int64, currency string) (*model.SpotAccount, error)
	GetAccountsByUserId(userId int64) ([]model.SpotAccount, error)
	HasBalance(userId int64, currency string, amount decimal.Decimal) (bool, error)
	GetLockedAccount(userId int64, currency string) (*model.SpotAccount, error)
	Freeze(req *model.FreezeReq) error
	Unfreeze(req *model.UnfreezeReq) error
}

type AccountService struct {
	spotAccountDao dao.ISpotAccountDao
}

func NewAccountService() *AccountService {
	spotAccountDao := dao.NewSpotAccountDao()
	return &AccountService{spotAccountDao: spotAccountDao}
}

func (s *AccountService) CreateAccount(userId int64, currency string) (bool, error) {
	return s.spotAccountDao.Create(userId, currency)
}

func (s *AccountService) CreateAccountList(userIds []int64, currencies []string) error {
	return s.spotAccountDao.CreateAccountList(userIds, currencies)
}

func (s *AccountService) GetExistsAccounts(userIds []int64, currency string) ([]int64, error) {
	return s.spotAccountDao.GetExistsAccounts(userIds, currency)
}

func (s *AccountService) GetAccount(userId int64, currency string) (*model.SpotAccount, error) {
	return s.spotAccountDao.GetAccount(userId, currency)
}

func (s *AccountService) GetAccountsByUserId(userId int64) ([]model.SpotAccount, error) {
	return s.spotAccountDao.GetAccountsByUserId(userId)
}

func (s *AccountService) HasBalance(userId int64, currency string, amount decimal.Decimal) (bool, error) {
	return s.spotAccountDao.HasBalance(userId, currency, amount)
}

func (s *AccountService) GetLockedAccount(userId int64, currency string) (*model.SpotAccount, error) {
	return s.spotAccountDao.GetLockedAccount(userId, currency)
}

func (s *AccountService) Freeze(req *model.FreezeReq) error {
	return s.spotAccountDao.Freeze(req)
}

func (s *AccountService) Unfreeze(req *model.UnfreezeReq) error {
	return s.spotAccountDao.Unfreeze(req)
}
