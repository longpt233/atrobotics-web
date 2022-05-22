package repository

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


type MySQLClient struct {
	db *gorm.DB
}


//TODO: singleton, close connection
func (myclient *MySQLClient) GetConn() *gorm.DB {

	conn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("ATRO_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"


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

func (myclient *MySQLClient) Close(){
	myclient.db.Close()
}

