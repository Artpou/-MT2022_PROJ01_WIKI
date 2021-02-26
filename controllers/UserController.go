package controllers

import (
	"net/http"

	"github.com/Artpou/wiki_golang/models"
	"github.com/Artpou/wiki_golang/views"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.AddUser(*user)))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.ShowUser(*user)))
}

func GetSelf(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.ShowUser(*user)))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.DeleteUser(*user)))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.DeleteUser(*user)))
}

func UpdateSelf(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.UpdateUser(*user)))
}

func DeleteSelf(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.DeleteUser(*user)))
}
