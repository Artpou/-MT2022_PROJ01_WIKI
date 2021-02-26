package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Artpou/wiki_golang/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/login/", login)
	router.HandleFunc("/api/comments/", comments)
	router.HandleFunc("/api/articles/", articles)

	/* User methods */
	router.HandleFunc("/api/users", ShowInfo).Methods("GET")
	router.HandleFunc("/api/users", UpdateInfo).Methods("POST")
	router.HandleFunc("/api/users", DeleteSelf).Methods("DELETE")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func ShowInfo(w http.ResponseWriter, r *http.Request) {
	controllers.GetSelf(w, r)
}

func UpdateInfo(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateSelf(w, r)
}

func DeleteSelf(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteSelf(w, r)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	// Création d'un utilisateur
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Success: User logged", id_user : "123"}`))
	}
}

func comments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	// Création d'un utilisateur
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Success: Add comment", id_user : "123"}`))

	case "GET":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Get comment", id_user : "123"}`))

	case "DELETE":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Sucess: Delete comment", id_user : "123"}`))

	case "PUT":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Sucess: Comment updated", id_user : "123"}`))

	}

}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	// Création d'un utilisateur
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Add article", id_user : "123"}`))

	case "GET":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "List of articles", id_user : "123"}`))

	case "DELETE":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Sucess: Delete article", id_user : "123"}`))

	case "PUT":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Sucess: Article updated", id_user : "123"}`))

	}

}
