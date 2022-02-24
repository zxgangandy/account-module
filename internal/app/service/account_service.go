package service

import (
	"account-module/internal/app/dao"
)

type IAccountService interface {
	CreateAccount(userId uint64, currency string) (bool, error)
	CreateAccountList(userIds []uint64, currencies []string) error
	GetExistAccounts(userIds []uint64, currency string) error
}

type AccountService struct {
	spotAccountDao dao.ISpotAccountDao
}

func NewAccountService() *AccountService {
	spotAccountDao := dao.NewSpotAccountDao()
	return &AccountService{spotAccountDao: spotAccountDao}
}

func (s *AccountService) CreateAccount(userId uint64, currency string) (bool, error) {
	return s.spotAccountDao.Create(userId, currency)
}

func (s *AccountService) CreateAccountList(userIds []uint64, currencies []string) error {
	return s.spotAccountDao.CreateAccountList(userIds, currencies)
}

func (s *AccountService) GetExistAccounts(userIds []uint64, currency string) ([]uint64, error) {
	return s.spotAccountDao.GetExistAccounts(userIds, currency)
}
