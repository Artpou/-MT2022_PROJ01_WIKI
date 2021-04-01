package router

import (
  "log"
  "fmt"
  "net/http"
  "github.com/Artpou/wiki_golang/controllers"
  "github.com/Artpou/wiki_golang/database"
  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
)

var db *gorm.DB

func HandleRequests() {
	log.Println("Starting development server at http://127.0.0.1:10000/")
	router := mux.NewRouter().StrictSlash(true)
  db = database.InitDb()
	//Login
	router.HandleFunc("/", HomeLink)
	router.HandleFunc("/api/login/", Login).Methods("POST")
	router.HandleFunc("/api/checkAuth/", CheckAuth).Methods("GET")
	router.HandleFunc("/api/checkAdmin/", CheckAdmin).Methods("GET")

	//Comments
	router.HandleFunc("/api/comments/", GetComments).Methods("GET")
	router.HandleFunc("/api/comments/", CreateComment).Methods("POST")
	router.HandleFunc("/api/comments/{id}", GetComment).Methods("GET")
	router.HandleFunc("/api/comments/{id}", UpdateComment).Methods("PUT")
	router.HandleFunc("/api/comments/{id}", DeleteComment).Methods("DELETE")

	//Articles
	router.HandleFunc("/api/articles/", GetArticles).Methods("GET")
	router.HandleFunc("/api/articles/", CreateArticle).Methods("POST")
	router.HandleFunc("/api/articles/{id}", GetArticle).Methods("GET")
	router.HandleFunc("/api/articles/{id}", UpdateArticle).Methods("PUT")
	router.HandleFunc("/api/articles/{id}", DeleteArticle).Methods("DELETE")

	//Users
	router.HandleFunc("/api/users/", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")

	//Self
	router.HandleFunc("/api/self/", GetSelf).Methods("GET")
	router.HandleFunc("/api/self/", UpdateSelf).Methods("PUT")
	router.HandleFunc("/api/self/", DeleteSelf).Methods("DELETE")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":10000", router))
}


//LOGIN

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func Login(w http.ResponseWriter, r *http.Request) {
	controllers.Signin(db, w, r)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	controllers.CheckAuth(w, r)
}

func CheckAdmin(w http.ResponseWriter, r *http.Request) {
	controllers.CheckAdmin(w, r)
}

//USER

func GetUsers(w http.ResponseWriter, r *http.Request) {
	controllers.GetUsers(db, w, r)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	controllers.GetUser(db, w, r)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateUser(db, w, r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	controllers.CreateUser(db, w, r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteUser(db, w, r)
}

func GetSelf(w http.ResponseWriter, r *http.Request) {
	controllers.GetSelf(db, w, r)
}

func UpdateSelf(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateSelf(db, w, r)
}

func DeleteSelf(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteSelf(db, w, r)
}

//ARTICLES

func GetArticles(w http.ResponseWriter, r *http.Request) {
	controllers.GetArticles(db, w, r)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	controllers.GetArticle(db, w, r)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	controllers.CreateArticle(db, w, r)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateArticle(db, w, r)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteArticle(db, w, r)
}

// COMMENTS

func GetComments(w http.ResponseWriter, r *http.Request) {
	controllers.GetComments(db, w, r)
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	controllers.GetComment(db, w, r)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	controllers.CreateComment(db, w, r)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateComment(db, w, r)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteComment(db, w, r)
}
