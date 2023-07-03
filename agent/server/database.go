package server

import (
	"easy-deploy/utils/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("deploy.db"), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}

	db.AutoMigrate(&types.WebProcess{})

	DB = db
}
