package yugle

import "github.com/gin-gonic/gin"

type Config struct {
	Mode          string
	Port          string
	DbDialect     string
	DataSourceUrl string
}

func DefaultConfig() *Config {
	config := Config{}
	config.Mode = gin.ReleaseMode
	config.Port = "8888"
	config.DbDialect = "mysql"
	config.DataSourceUrl = "root:sql123@/yugle?charset=utf8&parseTime=True&loc=Local"
	//config.DataSourceUrl = "root:sql123@tcp(127.0.0.1:3306)/yugle?charset=utf8&parseTime=True&loc=Local"
	return &config
}
