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

// ARTICLES

func GetArticles(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	articles := []models.Article{}
	db.Find(&articles)
	respond.RespondJSON(w, http.StatusOK, articles)
}

func GetArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	article := models.Article{}
	if err := db.First(&article, models.Article{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("Article"))
		return
	}

	comments := []models.Comment{}
	if err := db.Find(&comments, models.Comment{ArticleID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("Comment"))
		return
	}

	respond.RespondJSON(w, http.StatusOK, models.NewArticleWithComments(article, comments))
}

func CreateArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	claims, err := jwt.GetClaims(r)
	if !CheckAuth(w, r) || err != nil {
		return
	}

	rawArticle := models.Article{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawArticle); err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if rawArticle.Title == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("Title"))
		return
	}
	if rawArticle.Content == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("Content"))
		return
	}
	article := models.NewArticle(rawArticle.Title, rawArticle.Content, claims.ID)
	if err := db.Save(&article).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusCreated, article)
}

func UpdateArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAuth(w, r) {
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
	oldArticle := models.Article{}
	newArticle := models.Article{}
	if err := db.First(&oldArticle, models.Article{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("Article"))
		return
	}
	if !IsAdmin(w, r) {
		claims, err := jwt.GetClaims(r)
		if err != nil {
			respond.RespondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		authorID := claims.ID
		if oldArticle.AuthorID != authorID {
			respond.RespondError(w, http.StatusForbidden, "You don't have permission to update this article")
			return
		}
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newArticle); err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	updatedArticle := models.UpdateArticle(oldArticle, newArticle.Title, newArticle.Content)

	if err := db.Save(&updatedArticle).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusOK, updatedArticle)
}

func DeleteArticle(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAuth(w, r) {
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
	article := models.Article{}
	if err := db.First(&article, models.Article{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("Article"))
		return
	}
	if !IsAdmin(w, r) {
		claims, err := jwt.GetClaims(r)
		if err != nil {
			respond.RespondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		authorID := claims.ID
		if article.AuthorID != authorID {
			respond.RespondError(w, http.StatusForbidden, "You don't have permission to delete this article")
			return
		}
	}
	if err := db.Delete(&article).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusOK, views.DeleteArticle())
}
