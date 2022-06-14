package repository

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MySQLClient struct {
}

var lock = &sync.Mutex{}
var db *gorm.DB

//TODO: singleton, close connection
func (myclient *MySQLClient) GetConn() *gorm.DB {

	// db.LogMode(true) // show query log

	// db.AutoMigrate()
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {

			fmt.Println("Creating single instance now.")
			conn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("ATRO_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"

			var err error
			db, err = gorm.Open("mysql", conn)
			if err != nil {
				log.Fatal("Error connecting to Database" + err.Error())
				return nil
			} else {
				fmt.Println("Open conn successfully to " + conn)
			}

			// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
			db.DB().SetMaxIdleConns(10)

			// SetMaxOpenConns sets the maximum number of open connections to the database.
			db.DB().SetMaxOpenConns(100)

			// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
			db.DB().SetConnMaxLifetime(time.Hour)

		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return db
}

// func (myclient *MySQLClient) Close() {
// 	myclient.db.Close()
// }
