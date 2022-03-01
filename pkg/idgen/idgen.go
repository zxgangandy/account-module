package idgen

import (
	"account-module/pkg/conf"
	"account-module/pkg/datasource"
	gidConf "github.com/zxgangandy/gid/config"
	"strconv"
)

var (
	UidConfig *gidConf.DefaultUidConfig
)

func Init(conf *conf.AppConfig) {
	appPort := strconv.Itoa(conf.Application.Port)
	UidConfig = gidConf.New(datasource.GetDB(), appPort)
}

func GetIdgenConfig() *gidConf.DefaultUidConfig {
	return UidConfig
}
