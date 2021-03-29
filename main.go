package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Artpou/wiki_golang/controllers"
	"github.com/Artpou/wiki_golang/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	db = database.InitDb()
	handleRequests()
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

	// Users (à modifier)
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":10000", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	controllers.GetUsers(db, w, r)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	controllers.GetUser(db, w, r)
}

func updateUser(w http.ResponseWriter, r *http.Request){
	controllers.UpdateUser(db, w, r)
}

func createUser(w http.ResponseWriter, r *http.Request){
	controllers.CreateUser(db, w, r)
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	controllers.DeleteUser(db, w, r)
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

//ARTICLES

func getArticles(w http.ResponseWriter, r *http.Request) {
	controllers.GetArticles(db, w, r)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	controllers.GetArticle(db, w, r)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	controllers.CreateArticle(db, w, r)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateArticle(db, w, r)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteArticle(w, r)
}

// COMMENTS

func getComments(w http.ResponseWriter, r *http.Request) {
	controllers.GetComments(db, w, r)
}

func getComment(w http.ResponseWriter, r *http.Request) {
	controllers.GetComment(db, w, r)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	controllers.CreateComment(db, w, r)
}

func updateComment(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateComment(db, w, r)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteComment(db, w, r)
}
