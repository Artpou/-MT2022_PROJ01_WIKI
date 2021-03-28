package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Artpou/wiki_golang/models"
	"github.com/gorilla/mux"
)

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
	//Convert string to uint64
	u, err := strconv.ParseUint(key, 10, 64)
	if err == nil {
		//Convert uint64 to uint
		u := uint(u)
		article := db.First(&article, u)
		fmt.Println("Success : getting Article N.", key)
		fmt.Println(article)
		json.NewEncoder(w).Encode(article)
	} else {
		fmt.Println("Error : ID is not an integer", key)
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
	// return the string response containing the request body
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	json.Unmarshal(reqBody, &article)
	db.ModelUpdate(&article)
	fmt.Println("Success : updating Article N.", key)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Sucess: article deleted", id_user : "123"}`))
}
