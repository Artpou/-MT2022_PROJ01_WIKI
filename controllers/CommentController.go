package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Artpou/wiki_golang/handler"
	"github.com/Artpou/wiki_golang/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	comment := models.Comment{}
	if err := db.First(&comment, models.Comment{ID: uid}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, comment)
}

func GetComments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	comments := []models.Comment{}
	db.Find(&comments)
	handler.RespondJSON(w, http.StatusOK, comments)
}

func CreateComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !handler.IsAuthenticated(w, r) {
		return
	}
	rawComment := models.Comment{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawComment); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if rawComment.ArticleID == 0 {
		handler.RespondError(w, http.StatusBadRequest, "ArticleID is missing")
		return
	}
	if rawComment.Content == "" {
		handler.RespondError(w, http.StatusBadRequest, "Content is missing")
		return
	}
	comment := models.NewComment(rawComment.ArticleID, rawComment.Content)
	if err := db.Save(&comment).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, comment)
}

func UpdateComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !handler.IsAuthenticated(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	oldComment := models.Comment{}
	newComment := models.Comment{}
	if err := db.First(&oldComment, models.Comment{ID: uid}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newComment); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	updatedComment := models.UpdateComment(oldComment, newComment.Content)

	if err := db.Save(&updatedComment).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, updatedComment)
}

func DeleteComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !handler.IsAuthenticated(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	comment := models.Comment{}
	if err := db.First(&comment, models.Comment{ID: uid}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := db.Delete(&comment).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusNoContent, nil)
}
