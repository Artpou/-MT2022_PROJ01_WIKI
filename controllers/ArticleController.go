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
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	article := models.Article{}
	if err := db.First(&article, models.Article{ID: uid}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	comments := []models.Comment{}
	if err := db.Find(&comments, models.Comment{ArticleID: uid}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	handler.RespondJSON(w, http.StatusOK, models.ArticleWithComments{article, comments})
}

func CreateArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rawArticle := models.Article{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawArticle); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if rawArticle.Title == "" {
		handler.RespondError(w, http.StatusBadRequest, "Title is missing")
		return
	}
	if rawArticle.Content == "" {
		handler.RespondError(w, http.StatusBadRequest, "Content is missing")
		return
	}
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
