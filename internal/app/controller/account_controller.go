package controller

import (
	"account-module/internal/app/model"
	"account-module/internal/app/service"
	"account-module/internal/pkg/accounterr"
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
// @Router /v1/account [post]
func CreateAccount(c *gin.Context) {
	var req model.CreateAccountReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warnf("create account bind params err : %v", err)
		model.R.Error(c, baseerr.ErrBind.WithDetails(err.Error()))
		return
	}

	err := service.AccountServiceImpl.CreateAccount(req.UserID, req.Currency)
	if err != nil {
		model.R.Error(c, accounterr.ErrCreateAccount.WithDetails(err.Error()))
		return
	}

	model.R.Success(c, nil)
}
