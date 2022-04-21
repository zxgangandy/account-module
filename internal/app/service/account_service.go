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
	Deposit(req *model.DepositReq) error
	Withdraw(req *model.WithdrawReq) error
}

type AccountService struct {
	db             *gorm.DB
	spotAccountDao dao.ISpotAccountDao
	frozenDao      dao.IAccountFrozenDao
	unfreezeDao    dao.IAccountUnFreezeDao
	tradeDao       dao.IAccountTradeDao
	logDao         dao.IAccountLogDao
}

func NewAccountService() *AccountService {
	spotAccountDao := dao.NewSpotAccountDao()
	frozenDao := dao.NewAccountFrozenDao()
	unfreezeDao := dao.NewAccountUnfreezeDao()
	tradeDao := dao.NewAccountTradeDao()
	logDao := dao.NewAccountLogDao()
	return &AccountService{
		db:             datasource.GetDB(),
		spotAccountDao: spotAccountDao,
		frozenDao:      frozenDao,
		unfreezeDao:    unfreezeDao,
		tradeDao:       tradeDao,
		logDao:         logDao,
	}
}

func (s *AccountService) CreateAccount(userId int64, currency string) (bool, error) {
	return s.spotAccountDao.Save(userId, currency)
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

		err = s.frozenDao.Save(s.frozenDao.CreateFreezeOrder(account, req))
		if err != nil {
			return err
		}

		err = s.logDao.Save(s.logDao.CreateFreezeLog(account, req))
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *AccountService) Unfreeze(req *model.UnfreezeReq) error {
	return s.db.Transaction(func(db *gorm.DB) error {
		account, err := s.spotAccountDao.GetLockedAccount(req.UserId, req.Currency)
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

		err = s.unfreezeDao.Save(s.unfreezeDao.CreateUnfreezeOrder(frozen, req))
		if err != nil {
			return err
		}

		err = s.logDao.Save(s.logDao.CreateUnfreezeLog(account, req))
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *AccountService) Deposit(req *model.DepositReq) error {
	return s.db.Transaction(func(db *gorm.DB) error {
		account, err := s.spotAccountDao.GetLockedAccount(req.UserId, req.Currency)
		if err != nil {
			return err
		}

		result, err := s.spotAccountDao.DepositByUser(req)
		if !result {
			return errors.New("mysql error: deposit update account balance failed")
		} else if err != nil {
			return err
		}

		err = s.tradeDao.Save(s.tradeDao.CreateDepositOrder(account, req))
		if err != nil {
			return err
		}

		err = s.logDao.Save(s.logDao.CreateDepositLog(account, req))
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *AccountService) Withdraw(req *model.WithdrawReq) error {
	return s.db.Transaction(func(db *gorm.DB) error {
		account, err := s.spotAccountDao.GetLockedAccount(req.UserId, req.Currency)
		if err != nil {
			return err
		}

		if account.Balance.LessThan(req.Amount) {
			return errors.Errorf("withdraw amount: %v bigger than balance "+
				"amount: %v", req.Amount, account.Balance)
		}

		result, err := s.spotAccountDao.WithdrawByUser(req)
		if !result {
			return errors.Errorf("withdraw balance is not enough, "+
				"balance=%v, withdraw amount=%v", account.Balance, req.Amount)
		} else if err != nil {
			return err
		}

		err = s.tradeDao.Save(s.tradeDao.CreateWithdrawOrder(account, req))
		if err != nil {
			return err
		}

		err = s.logDao.Save(s.logDao.CreateWithdrawLog(account, req))
		if err != nil {
			return err
		}

		return nil
	})
}
