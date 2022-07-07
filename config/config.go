package config

import (
	"Project_Rent_Book/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/Project_Rent_Book?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Book{})
	db.AutoMigrate(&entity.PinjamBuku{})
	//db.Set(le_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	return db
}
