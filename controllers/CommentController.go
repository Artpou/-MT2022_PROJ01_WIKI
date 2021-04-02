package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Artpou/wiki_golang/handler/jwt"
	"github.com/Artpou/wiki_golang/handler/respond"
	"github.com/Artpou/wiki_golang/models"
	"github.com/Artpou/wiki_golang/views"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	comment := models.Comment{}
	if err := db.First(&comment, models.Comment{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("Comment"))
		return
	}
	respond.RespondJSON(w, http.StatusOK, comment)
}

func GetComments(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	comments := []models.Comment{}
	db.Find(&comments)
	respond.RespondJSON(w, http.StatusOK, comments)
}

func CreateComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(w, r) {
		return
	}
	rawComment := models.Comment{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawComment); err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if rawComment.ArticleID == 0 {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("ArticleID"))
		return
	}
	if rawComment.Content == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("Content"))
		return
	}
	comment := models.NewComment(rawComment.ArticleID, rawComment.Content)
	if err := db.Save(&comment).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusCreated, comment)
}

func UpdateComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	oldComment := models.Comment{}
	newComment := models.Comment{}
	if err := db.First(&oldComment, models.Comment{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("Comment"))
		return
	}
	if !IsAdmin(w, r) {
		claims, err := jwt.GetClaims(r)
		if err != nil {
			respond.RespondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		authorID := claims.ID
		if oldComment.AuthorID != authorID {
			respond.RespondError(w, http.StatusForbidden, "You don't have permission to update this comment")
			return
		}
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newComment); err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	updatedComment := models.UpdateComment(oldComment, newComment.Content)

	if err := db.Save(&updatedComment).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusOK, updatedComment)
}

func DeleteComment(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	comment := models.Comment{}
	if err := db.First(&comment, models.Comment{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("Comment"))
		return
	}
	if !IsAdmin(w, r) {
		claims, err := jwt.GetClaims(r)
		if err != nil {
			respond.RespondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		authorID := claims.ID
		if comment.AuthorID != authorID {
			respond.RespondError(w, http.StatusForbidden, "You don't have permission to delete this comment")
			return
		}
	}
	if err := db.Delete(&comment).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusNoContent, nil)
}
