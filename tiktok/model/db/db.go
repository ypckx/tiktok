package db

import (
	"fmt"
	"tiktok/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dB *gorm.DB

func InitDB() {
	var err error
	dB, err = gorm.Open(mysql.Open(config.Dns),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	table_options := "ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci"

	// AutoMigrate函数可以快速建表，如果表已经存在不会重复创建
	/*
		解释为什么要分三次建表：因为Video表中有外键依赖于User表。引用不存在的表，无法创建外键
	*/
	err = dB.Set("gorm:table_options", table_options).
		AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	err = dB.Set("gorm:table_options", table_options).
		AutoMigrate(&Video{})
	if err != nil {
		panic(err)
	}

	err = dB.Set("gorm:table_options", table_options).
		AutoMigrate(&Favorite{}, &Relation{}, &Comment{})
	if err != nil {
		panic(err)
	}

	fmt.Println("connect database OK!")
}

func GetDB() *gorm.DB {
	return dB
}

func CloseDB() {

}
