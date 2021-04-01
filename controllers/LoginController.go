package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Artpou/wiki_golang/handler"
	"github.com/Artpou/wiki_golang/models"
	"github.com/jinzhu/gorm"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func Signin(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{}
	if err := db.First(&user, models.User{Username: creds.Username}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, "Username doesn't exist")
		return
	}

	if !handler.CheckPasswordHash(creds.Password, user.Password) {
		handler.RespondError(w, http.StatusUnauthorized, "Wrong password")
		return
	}

	token, err := handler.SetToken(creds.Username, w)

	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	handler.RespondJSON(w, http.StatusCreated, token)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	if handler.IsAuthenticated(w, r) {
		handler.RespondJSON(w, http.StatusFound, "You are authenticated !")
	}
}
