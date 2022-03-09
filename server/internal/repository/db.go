package repository

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//TODO: singleton, close connection
func DB() *gorm.DB {

	conn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("ATRO_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Fatal("Error connecting to Database" + err.Error())
		return nil
	} else {
		fmt.Println("Open conn successfully to " + conn)
	}

	db.LogMode(true) // show query log
	db.AutoMigrate()
	return db
}