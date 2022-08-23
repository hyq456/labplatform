package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"labplatform/utils"
	"os"
	"time"
)

// 全局mysql数据库变量
var db *gorm.DB
var err error

//var dns = utils.DbUser + ":" + utils.DbPassWord + "@tcp(" + utils.DbHost + ":" + utils.DbPort + ")" +
//	"/" + utils.DbName + "?charset=utf8&parseTime=True&loc=Local"

func InitDb() {
	//fmt.Println(dsn)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}

	db.AutoMigrate(&User{})

	sqlDB, err := db.DB()
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
