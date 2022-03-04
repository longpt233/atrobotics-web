package repository

import (
	"atro/internal/model"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DB -> connection to Database
func DB() *gorm.DB {

	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Error connecting to Database" + err.Error())
		return nil
	}
	db.LogMode(true)

	db.AutoMigrate(&model.User{})
	return db

}
