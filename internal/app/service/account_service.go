package service

import (
	"account-module/internal/app/dao"
	"container/list"
)

type IAccountService interface {
	CreateAccount(userId int64, currency string) error
	CreateAccountList(userIds *list.List, currencies *list.List) error
}

type AccountService struct {
	spotAccountDao dao.ISpotAccountDao
}

func NewAccountService() *AccountService {
	spotAccountDao := dao.NewSpotAccountDao()
	return &AccountService{spotAccountDao: spotAccountDao}
}

func (s *AccountService) CreateAccount(userId int64, currency string) error {
	return s.spotAccountDao.Create(userId, currency)
}

func (s *AccountService) CreateAccountList(userIds *list.List, currencies *list.List) error {
	return nil
}
