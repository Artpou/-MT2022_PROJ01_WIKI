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
  articles := []models.Article{}
  db.Find(&articles)
  json.NewEncoder(w).Encode(articles)
	//w.Write([]byte(fmt.Sprintf(`{"%v"}`, articles)))
}

func GetArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	u, err := strconv.ParseUint(key, 10, 64)
	if err == nil {
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
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	json.Unmarshal(reqBody, &article)
	db.Create(&article)
	w.Write([]byte("Success : Creating Article"))
}

func UpdateArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  /*var title, content string
  vars := mux.Vars(r)
	key := vars["id"]
  u, err := strconv.ParseUint(key, 10, 64)
	if err == nil {
    u := uint(u)
    article := db.First(models.Article{}, u)
  	reqBody, _ := ioutil.ReadAll(r.Body)
    db.Model(&article).Update()
  	fmt.Println("Success : updating Article N.", key)
  	json.NewEncoder(w).Encode(article)
  }*/
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Sucess: article deleted", id_user : "123"}`))
}
