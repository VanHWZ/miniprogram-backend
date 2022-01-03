package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"miniprogram-backend/conf"
)

var DB *gorm.DB

func init() {
	dbConf := conf.Config.Database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbConf.Host, dbConf.User, dbConf.Password, dbConf.Dbname, dbConf.Port, dbConf.Sslmode, dbConf.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	db.Logger.LogMode(logger.Info)
	if err != nil {
		fmt.Println(err)
	} else {
		db.AutoMigrate(&Group{}, &User{}, &Message{}, &Event{})
	}
	DB = db
}
