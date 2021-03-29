package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Artpou/wiki_golang/handler"
	"github.com/Artpou/wiki_golang/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// ARTICLES

func GetArticles(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	articles := []models.Article{}
	db.Find(&articles)
	fmt.Println(articles)
	handler.RespondJSON(w, http.StatusOK, articles)
}

func GetArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	u, err := strconv.ParseUint(key, 10, 64)
	if err == nil {
		u := uint(u)
		article := db.First(models.Article{}, u)
		fmt.Println("Success : getting Article N.", key)
		handler.RespondJSON(w, http.StatusOK, article)
	} else {
		fmt.Println("Error : ID is not an integer", key)
	}
}

func CreateArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rawArticle := models.Article{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawArticle); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	article := models.NewArticle(rawArticle.Title, rawArticle.Content)
	if err := db.Save(&article).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, article)
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
