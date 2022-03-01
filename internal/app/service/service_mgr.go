package service

import (
	"account-module/pkg/conf"
	"account-module/pkg/idgen"
)

var (
	AccountServiceImpl IAccountService
)

func Init() {
	initServices()
	initIdGen()
}

func initServices() {
	AccountServiceImpl = NewAccountService()
}

func initIdGen() {
	idgen.Init(conf.Config)
}
