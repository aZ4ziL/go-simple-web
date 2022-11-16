package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Deklarasi variabel gorm.DB
var db *gorm.DB

func init() {
	db = GetDB()

	// Migrate model
	db.AutoMigrate(&Test{})
}

// Buat koneksi database
func GetDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/go-simple-web?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
