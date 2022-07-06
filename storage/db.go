package storage

import (
	"API-Golang-WithDB/config"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMySQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(mysql.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return DB
}
