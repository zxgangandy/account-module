package idgen

import (
	"account-module/pkg/conf"
	"account-module/pkg/datasource"
	"github.com/zxgangandy/gid"
	gidConf "github.com/zxgangandy/gid/config"
	"strconv"
)

var (
	IdGenerator *gid.DefaultUidGenerator
)

func Init(conf *conf.AppConfig) {
	appPort := strconv.Itoa(conf.Application.Port)
	config := gidConf.New(datasource.GetDB(), appPort)
	IdGenerator = gid.New(config)
}

func Get() *gid.DefaultUidGenerator {
	return IdGenerator
}
