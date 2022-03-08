package main

import (
	"atro/internal/router"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	} else {
		log.Println("Load .env succesfullyyy")
	}
}

func main() {

	fmt.Println("Main Application Starts")
	//Loading Environmental Variable
	loadenv()
	log.Fatal(router.RunAPI(":8091"))
}
