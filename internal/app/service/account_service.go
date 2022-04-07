package service

import (
	"account-module/internal/app/dao"
	"account-module/internal/app/model"
)

type IAccountService interface {
	CreateAccount(userId int64, currency string) (bool, error)
	CreateAccountList(userIds []int64, currencies []string) error
	GetExistsAccounts(userIds []int64, currency string) ([]int64, error)
	GetAccount(userId int64, currency string) (model.SpotAccount, error)
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

func (s *AccountService) GetAccount(userId int64, currency string) (model.SpotAccount, error) {
	return s.spotAccountDao.GetAccount(userId, currency)
}
