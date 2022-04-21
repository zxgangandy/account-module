package controller

import (
	"account-module/internal/app/intererr"
	"account-module/internal/app/model"
	"account-module/internal/app/service"
	"account-module/pkg/baseerr"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

// CreateAccount 创建用户账户
// @Summary 通过用户id和币种创建用户账户
// @Description Create an account by user id and currency
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/create_one [post]
func CreateAccount(c *gin.Context) {
	var req model.CreateAccountReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("create account bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	result, err := service.AccountServiceImpl.CreateAccount(req.UserID, req.Currency)
	if err != nil {
		model.R.Error(c, intererr.ErrCreateAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, result)
}

// CreateAccount 创建用户账户列表
// @Summary 通过用户id和币种创建用户账户列表
// @Description Create an account by user id list and currency list
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/create_list [post]
func CreateAccounts(c *gin.Context) {
	var req model.CreateAccountListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("create account bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	err := service.AccountServiceImpl.CreateAccountList(req.UserIDList, req.CurrencyList)
	if err != nil {
		logger.Errorf("create account list err : %v", err)
		model.R.Error(c, intererr.ErrCreateAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, nil)
}

// CreateAccount 创建用户账户列表
// @Summary 通过用户id列表和币种获取已经创建、存在的用户账户列表
// @Description get accounts by user id list and the currency
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/exist_list [post]
func GetExistsAccounts(c *gin.Context) {
	var req model.ExistAccountListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("Get existing account bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	userIds, err := service.AccountServiceImpl.GetExistsAccounts(req.UserIDList, req.Currency)
	if err != nil {
		logger.Errorf("Get existing account list err : %v", err)
		model.R.Error(c, intererr.ErrGetExitsAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, userIds)
}

// FindAccount 查询某个用户的某币种账户
// @Summary 通过用户id列表和币种获取已经创建、存在的用户账户
// @Description get account by user id and the currency
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/get_one [post]
func FindAccount(c *gin.Context) {
	var req model.GetAccountReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("Get account bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	account, err := service.AccountServiceImpl.GetAccount(req.UserId, req.Currency)
	if err != nil {
		logger.Errorf("Get account err : %v", err)
		model.R.Error(c, intererr.ErrGetExitsAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, account)
}

// FindAccounts 查询用户账户列表
// @Summary 通过用户id列表已经创建、存在的用户账户列表
// @Description get account by user id
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/find_list [post]
func FindAccounts(c *gin.Context) {
	var req model.GetAccountsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("Find accounts bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	accounts, err := service.AccountServiceImpl.GetAccountsByUserId(req.UserId)
	if err != nil {
		logger.Errorf("Find accounts err : %v", err)
		model.R.Error(c, intererr.ErrGetExitsAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, accounts)
}

// HasBalance 查询用户账户余额是否足够
// @Summary 通过用户id、币种查询用户账户余额是否足够
// @Description check用户余额
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/has_balance [post]
func HasBalance(c *gin.Context) {
	var req model.HasBalanceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("Has balance bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	userIds, err := service.AccountServiceImpl.HasBalance(req.UserId, req.Currency, req.Amount)
	if err != nil {
		logger.Errorf("Has balance err : %v", err)
		model.R.Error(c, intererr.ErrGetExitsAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, userIds)
}

// Frozen 冻结用户账户余额
// @Summary 通过用户id、币种冻结用户余额
// @Description 冻结用户余额
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/freeze_balance [post]
func Freeze(c *gin.Context) {
	var req model.FreezeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("Frozen bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	err := service.AccountServiceImpl.Freeze(&req)
	if err != nil {
		logger.Errorf("Frozen err : %v", err)
		model.R.Error(c, intererr.ErrGetExitsAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, true)
}

// Frozen 冻结用户账户余额
// @Summary 通过用户id、币种冻结用户余额
// @Description 冻结用户余额
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/unfreeze_balance [post]
func Unfreeze(c *gin.Context) {
	var req model.UnfreezeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("Frozen bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	err := service.AccountServiceImpl.Unfreeze(&req)
	if err != nil {
		logger.Errorf("Frozen err : %v", err)
		model.R.Error(c, intererr.ErrGetExitsAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, true)
}

// Frozen 冻结用户账户余额
// @Summary 通过用户id、币种冻结用户余额
// @Description 冻结用户余额
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/deposit [post]
func Deposit(c *gin.Context) {
	var req model.DepositReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("Deposit bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	err := service.AccountServiceImpl.Deposit(&req)
	if err != nil {
		logger.Errorf("Deposit err : %v", err)
		model.R.Error(c, intererr.ErrGetExitsAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, true)
}

// Frozen 冻结用户账户余额
// @Summary 通过用户id、币种冻结用户余额
// @Description 冻结用户余额
// @Tags 账户
// @Accept  json
// @Produce  json
// @Param
// @Success 200 {object}
// @Router /v1/account/withdraw [post]
func Withdraw(c *gin.Context) {
	var req model.WithdrawReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("Withdraw bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	err := service.AccountServiceImpl.Withdraw(&req)
	if err != nil {
		logger.Errorf("Withdraw err : %v", err)
		model.R.Error(c, intererr.ErrGetExitsAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, true)
}
