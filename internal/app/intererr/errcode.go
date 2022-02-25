package intererr

import "account-module/pkg/baseerr"

var (
	// account errors
	ErrCreateAccount   = baseerr.NewError(20101, "创建账户错误")
	ErrGetExitsAccount = baseerr.NewError(20102, "获取存在的账户错误")
)
