package service

import (
	"account-module/internal/app/dao"
	"account-module/internal/app/model"
	"account-module/pkg/datasource"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type IAccountService interface {
	CreateAccount(userId int64, currency string) (bool, error)
	CreateAccountList(userIds []int64, currencies []string) error
	GetExistsAccounts(userIds []int64, currency string) ([]int64, error)
	GetAccount(userId int64, currency string) (*model.SpotAccount, error)
	GetAccountsByUserId(userId int64) ([]model.SpotAccount, error)
	HasBalance(userId int64, currency string, amount decimal.Decimal) (bool, error)
	Freeze(req *model.FreezeReq) error
	Unfreeze(req *model.UnfreezeReq) error
}

type AccountService struct {
	db             *gorm.DB
	spotAccountDao dao.ISpotAccountDao
	frozenDao      dao.IAccountFrozenDao
	unfreezeDao    dao.IAccountUnFreezeDao
	logDao         dao.IAccountLogDao
}

func NewAccountService() *AccountService {
	spotAccountDao := dao.NewSpotAccountDao()
	frozenDao := dao.NewAccountFrozenDao()
	unfreezeDao := dao.NewAccountUnfreezeDao()
	logDao := dao.NewAccountLogDao()
	return &AccountService{
		db:             datasource.GetDB(),
		spotAccountDao: spotAccountDao,
		frozenDao:      frozenDao,
		unfreezeDao:    unfreezeDao,
		logDao:         logDao,
	}
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

func (s *AccountService) Freeze(req *model.FreezeReq) error {
	return s.db.Transaction(func(db *gorm.DB) error {
		account, err := s.spotAccountDao.GetLockedAccount(req.UserId, req.Currency)
		if err != nil {
			return err
		}

		result, err := s.spotAccountDao.FreezeByUser(req)
		if !result {
			return errors.Errorf("freeze failed, balance is not enough, "+
				"balance=%v, freeze amount=%v", account.Balance, req.Amount)
		} else if err != nil {
			return err
		}

		err = s.frozenDao.Create(s.spotAccountDao.CreateFreezeOrder(account, req))
		if err != nil {
			return err
		}

		err = s.logDao.Create(s.spotAccountDao.CreateFreezeLog(account, req))
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *AccountService) Unfreeze(req *model.UnfreezeReq) error {
	return s.db.Transaction(func(db *gorm.DB) error {
		account, err := s.GetLockedAccount(req.UserId, req.Currency)
		if err != nil {
			return err
		}

		if account.Frozen.LessThan(req.Amount) {
			return errors.Errorf("unfreeze amount: %v bigger than frozen "+
				"amount: %v", req.Amount, account.Frozen)
		}

		result, err := s.spotAccountDao.UnfreezeByUser(req)
		if !result {
			return errors.New("unfreeze failed: maybe balance not enough")
		} else if err != nil {
			return err
		}

		result, err = s.frozenDao.UpdateUnfreeze(req)
		if !result {
			return errors.New("mysql error: update freeze order failed")
		} else if err != nil {
			return err
		}

		frozen, err := s.frozenDao.Get(req.OrderId, req.BizType)
		if err != nil {
			return err
		}

		err = s.unfreezeDao.Create(s.spotAccountDao.CreateUnfreezeOrder(frozen, req))
		if err != nil {
			return err
		}

		err = s.logDao.Create(s.spotAccountDao.CreateUnfreezeLog(account, req))
		if err != nil {
			return err
		}

		return nil
	})
}
