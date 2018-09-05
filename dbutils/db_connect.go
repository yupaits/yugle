package dbutils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"yugle/config"
)

func Connect() *gorm.DB {
	appConfig := config.DefaultConfig()
	db, err := gorm.Open(appConfig.DbDialect, appConfig.DataSourceUrl)
	if err != nil {
		log.Panic()
		return nil
	}
	return db
}
