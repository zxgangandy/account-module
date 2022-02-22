package accounterr

import "account-module/pkg/baseerr"

var (
	// account errors
	ErrCreateAccount = baseerr.NewError(20101, "创建账户失败")
)
