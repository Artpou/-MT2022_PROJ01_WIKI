package database

import (
	"log"

	"github.com/Artpou/wiki_golang/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error
var dbServer, dbName, dbUsername, dbPassword, dbPort, dbConn string

func InitDb() *gorm.DB {
	//init bdd
	dbServer = "sql11.freemysqlhosting.net"
	dbName = "sql11395463"
	dbUsername = "sql11395463"
	dbPassword = "5mRSPiqM9M"
	dbPort = "3306"
	dbConn = dbUsername + ":" + dbPassword + "@tcp(" + dbServer + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True"

	db, err = gorm.Open("mysql", dbConn)

	if err != nil {
		log.Println("DB connection Failed to Open")
	} else {
		log.Println("DB connection Established")
	}

	db.AutoMigrate(&models.Comment{}, &models.Article{}, &models.User{})

	return db
}

func GetDb() *gorm.DB {
	return db
}
