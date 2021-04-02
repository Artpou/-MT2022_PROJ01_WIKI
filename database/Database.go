package database

import (
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/Artpou/wiki_golang/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type dbConfig struct {
	dbServer  	string
	dbName    	string
	dbUsername  string
	dbPassword  string
	dbPort  	  string
}
var db *gorm.DB
var err error

func InitDb() *gorm.DB {
	config := dbConfig{}
	jsonFile, _ := os.Open("config.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &config)
	dbConn := config.dbUsername + ":" + config.dbPassword + "@tcp(" + config.dbServer + ":" + config.dbPort + ")/" + config.dbName + "?charset=utf8&parseTime=True"
	db, err = gorm.Open("mysql", dbConn)

	if err != nil {
		log.Println("DB connection Failed to Open")
	} else {
		log.Println("DB connection Established")
		db.AutoMigrate(&models.Comment{}, &models.Article{}, &models.User{})
	}
	return db
}

func GetDb() *gorm.DB {
	return db
}
