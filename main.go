package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Artpou/wiki_golang/controllers"
	"github.com/Artpou/wiki_golang/models"
)

var db *gorm.DB
var err error
var dbServer, dbName, dbUsername, dbPassword, dbPort, dbConn string


func main() {
	//init bdd
	dbServer   = "sql11.freemysqlhosting.net"
  dbName     = "sql11395463"
	dbUsername = "sql11395463"
	dbPassword = "5mRSPiqM9M"
	dbPort     = "3306"
	dbConn		 = dbUsername+":"+dbPassword+"@tcp("+dbServer+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True"

	db, err		 = gorm.Open("mysql", dbConn)

	if err != nil {
  	log.Println("DB connection Failed to Open")
  } else {
  	log.Println("DB connection Established")
  }

	db.AutoMigrate(&models.Comment{}, &models.Article{}, &models.User{} )
	handleRequests();
	defer db.Close()
}

func handleRequests(){
	log.Println("Starting development server at http://127.0.0.1:10000/")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/login/", login).Methods("POST")

	//Comments
	router.HandleFunc("/api/comments/", controllers.getComments()).Methods("GET")
	router.HandleFunc("/api/comments/", controllers.createComment()).Methods("POST")
	router.HandleFunc("/api/comments/{id}", controllers.getComment()).Methods("GET")
	router.HandleFunc("/api/comments/{id}", controllers.updateComment()).Methods("PUT")
	router.HandleFunc("/api/comments/{id}", controllers.deleteComment()).Methods("DELETE")

	//Articles
	router.HandleFunc("/api/articles/", getArticles).Methods("GET")
	router.HandleFunc("/api/articles/", createArticle).Methods("POST")
	router.HandleFunc("/api/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/api/articles/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/api/articles/{id}", deleteArticle).Methods("DELETE")

	// User methods
	router.HandleFunc("/api/users", showInfo).Methods("GET")
	router.HandleFunc("/api/users", updateInfo).Methods("POST")
	router.HandleFunc("/api/users", deleteSelf).Methods("DELETE")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":10000", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func showInfo(w http.ResponseWriter, r *http.Request) {
	controllers.GetSelf(w, r)
}

func updateInfo(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateSelf(w, r)
}

func deleteSelf(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteSelf(w, r)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Success: User logged", id_user : "123"}`))
}
