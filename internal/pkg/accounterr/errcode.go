package accounterr

import "account-module/pkg/baseerr"

var (
	// user errors
	ErrCreateAccount         = baseerr.NewError(20101, "创建账户失败")
	ErrPasswordIncorrect     = baseerr.NewError(20102, "账号或密码错误")
	ErrAreaCodeEmpty         = baseerr.NewError(20103, "手机区号不能为空")
	ErrPhoneEmpty            = baseerr.NewError(20104, "手机号不能为空")
	ErrGenVCode              = baseerr.NewError(20105, "生成验证码错误")
	ErrSendSMS               = baseerr.NewError(20106, "发送短信错误")
	ErrSendSMSTooMany        = baseerr.NewError(20107, "已超出当日限制，请明天再试")
	ErrVerifyCode            = baseerr.NewError(20108, "验证码错误")
	ErrEmailOrPassword       = baseerr.NewError(20109, "邮箱或密码错误")
	ErrTwicePasswordNotMatch = baseerr.NewError(20110, "两次密码输入不一致")
	ErrRegisterFailed        = baseerr.NewError(20111, "注册失败")
)
