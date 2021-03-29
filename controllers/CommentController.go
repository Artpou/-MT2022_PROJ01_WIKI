package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Artpou/wiki_golang/handler"
	"github.com/Artpou/wiki_golang/models"
	"github.com/jinzhu/gorm"
)

func GetComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func GetComments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	comments := []models.Comment{}
	db.Find(&comments)
	fmt.Println(comments)
	handler.RespondJSON(w, http.StatusOK, comments)
}

func CreateComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rawComment := models.Comment{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawComment); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	comment := models.NewComment(rawComment.ArticleID, rawComment.Content)
	if err := db.Save(&comment).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, comment)
}

func DeleteComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func UpdateComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}
