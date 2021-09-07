package service

var (
	AccountServiceImpl IAccountService
)

func Init() {
	AccountServiceImpl = NewAccountService()
}
