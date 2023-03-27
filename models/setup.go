package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=35.187.248.198 user=postgres password=d3v3l0p8015 dbname=Tral_Week1_Vincent port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic("failed to get sql.db from gorm.db")
		}
		sqlDB.Close()
	}()

	database.AutoMigrate(&Post{})
	DB = database
}
