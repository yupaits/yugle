package yugle

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func DbConnect() *gorm.DB {
	appConfig := DefaultConfig()
	db, err := gorm.Open(appConfig.DbDialect, appConfig.DataSourceUrl)
	if err != nil {
		log.Panic()
		return nil
	}
	return db
}
