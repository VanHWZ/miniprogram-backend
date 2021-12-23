package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DatabaseInit() {
	configString := "host=localhost user=vincent password=990130 dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(configString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	db.Logger.LogMode(logger.Info)
	if err != nil {
		fmt.Println(err)
	} else {
		db.AutoMigrate(&User{}, &Message{}, &Event{}, &Group{})
	}
	DB = db
}
