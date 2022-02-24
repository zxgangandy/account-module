package service

import (
	"account-module/internal/app/dao"
)

type IAccountService interface {
	CreateAccount(userId uint64, currency string) (bool, error)
	CreateAccountList(userIds []uint64, currencies []string) error
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
