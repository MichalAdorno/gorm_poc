package main

import (
	"gorm_poc/connector"
	"gorm_poc/router"
	"log"
	"net/http"

	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// var db *gorm.DB

func main() {
	dbConfig := connector.ReadInDbConfig()
	connString := dbConfig.GetConnectionString()
	log.Println(connString)
	// db, err := gorm.Open("postgres", dbConfig.GetConnectionString())
	log.Fatal(http.ListenAndServe(":8080", router.InitRouter()))
}
