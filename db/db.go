package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"labplatform/model"
	"labplatform/utils"
	"time"
)

// 全局mysql数据库变量
var DB *gorm.DB
var err error

var dsn = utils.DbUser + ":" + utils.DbPassWord + "@tcp(" + utils.DbHost + ":" + utils.DbPort + ")" +
	"/" + utils.DbName + "?charset=utf8&parseTime=True&loc=Local"

func InitDb() {
	//fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败")
	}

	DB.AutoMigrate(&model.User{})

	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("打开数据库失败")
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
