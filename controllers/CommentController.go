package controllers

import (
	"net/http"

	"github.com/Artpou/wiki_golang/models"
	"github.com/Artpou/wiki_golang/views"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	comment := models.NewComment("test")
	w.Write([]byte(views.AddComment(*comment)))
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	comment := models.NewComment("test")
	w.Write([]byte(views.GetComment(*comment)))
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(views.GetComments()))
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	user := models.NewComment("test")
	w.Write([]byte(views.DeleteComment(*user)))
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	user := models.NewComment("test")
	w.Write([]byte(views.UpdateComment(*user)))
}
