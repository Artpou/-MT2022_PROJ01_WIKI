package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Artpou/wiki_golang/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// ARTICLES

func GetArticles(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	result := db.First(&models.Article{})
	fmt.Println("YES")

	//json.NewEncoder(w).Encode(articles)
	w.Write([]byte(fmt.Sprintf(`{"%v"}`, result)))
}

func GetArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	//Convert string to uint64
	u, err := strconv.ParseUint(key, 10, 64)
	if err == nil {
		//Convert uint64 to uint
		u := uint(u)
		article := db.First(models.Article{}, u)

		fmt.Println("Success : getting Article N.", key)
		fmt.Println(article)
		json.NewEncoder(w).Encode(article)
	} else {
		fmt.Println("Error : ID is not an integer", key)
	}
}

func CreateArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	//return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	json.Unmarshal(reqBody, &article)
	db.Create(&article)
	w.Write([]byte("Success : Creating Article"))
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	// return the string response containing the request body
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	json.Unmarshal(reqBody, &article)
	//db.ModelUpdate(&article)
	fmt.Println("Success : updating Article N.", key)
	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Sucess: article deleted", id_user : "123"}`))
}
