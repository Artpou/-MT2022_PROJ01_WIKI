package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Artpou/wiki_golang/controllers"
	"github.com/Artpou/wiki_golang/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error
var dbServer, dbName, dbUsername, dbPassword, dbPort, dbConn string

func main() {
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
	handleRequests()
	defer db.Close()
}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:10000/")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/login/", login).Methods("POST")

	//Comments
	router.HandleFunc("/api/comments/", getComments).Methods("GET")
	router.HandleFunc("/api/comments/", createComment).Methods("POST")
	router.HandleFunc("/api/comments/{id}", getComment).Methods("GET")
	router.HandleFunc("/api/comments/{id}", updateComment).Methods("PUT")
	router.HandleFunc("/api/comments/{id}", deleteComment).Methods("DELETE")

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

// ARTICLES

func getArticles(w http.ResponseWriter, r *http.Request) {
	articles := []models.Article{}
	db.Find(&articles)
	fmt.Println("Success : getting all articles")
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	articles := []models.Article{}
	db.Find(&articles)
	for _, article := range articles {
		//Convert string to uint64
		u, err := strconv.ParseUint(key, 10, 64)
		if err == nil {
			//Convert uint64 to uint
			u := uint(u)
			if article.ID == u {
				fmt.Println(article)
				fmt.Println("Success : getting Article N.", key)
				json.NewEncoder(w).Encode(article)
			}
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	json.Unmarshal(reqBody, &article)
	db.Create(&article)
	fmt.Println("Success : Creating Article")
	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Sucess: article updated", id_user : "123"}`))
	controllers.DeleteSelf(w, r)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Sucess: article deleted", id_user : "123"}`))
	controllers.DeleteSelf(w, r)

}

// COMMENTS

func getComments(w http.ResponseWriter, r *http.Request) {
	controllers.GetComments(w, r)
}

func getComment(w http.ResponseWriter, r *http.Request) {
	controllers.GetComment(w, r)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	controllers.AddComment(w, r)
}

func updateComment(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateComment(w, r)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteComment(w, r)
}
