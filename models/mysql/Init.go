package mysql

import (
	"GinWeb/common/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDb *gorm.DB
/**
mysql建立连接
 */
func InitConnect() {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Name,
		config.PASSWORD,
		config.HOST,
		config.PORT,
		config.DATABASE)
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {

		panic("failed to connect database")
	}
	MysqlDb = db
	//return db

}
