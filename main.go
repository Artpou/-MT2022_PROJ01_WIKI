package main

import (
	//"fmt"
	"net/http"
	//"log"
)

func test() {

	http.HandleFunc("/api/login/", login)
	http.HandleFunc("/api/users/", users)
	http.HandleFunc("/api/comments/", comments)
	http.HandleFunc("/api/articles/", articles)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
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

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	// Création d'un utilisateur
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Success: User created", id_user : "123"}`))

	case "GET":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "SHOW INFO", id_user : "123"}`))

	case "DELETE":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Sucess: Users deleted", id_user : "123"}`))

	case "PUT":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Sucess: Users updated", id_user : "123"}`))

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
