package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type AppConfig struct {
	DataSource  *DatasourceConfig
	Application *Application
}

type DatasourceConfig struct {
	DriverName string
	Addr       string
	Database   string
	User       string
	Password   string
	Charset    string
}

type Application struct {
	Name          string
	Mode          string
	Port          int
	LogLevel      string
	LogOutPath    string
	LogMaxSaveDay int
}

func GetConfig() *AppConfig {
	var appConfig AppConfig
	var configOnce sync.Once

	configOnce.Do(func() {
		config := viper.New()
		config.SetConfigName("application.yml")
		config.AddConfigPath("conf")
		config.SetConfigType("yaml")
		err := config.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		if err := config.Unmarshal(&appConfig); err != nil {
			panic(fmt.Errorf("read config file to struct err: %s \n", err))
		}
	})

	return &appConfig
}
