package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Artpou/wiki_golang/handler/jwt"
	"github.com/Artpou/wiki_golang/handler/password"
	"github.com/Artpou/wiki_golang/handler/respond"
	"github.com/Artpou/wiki_golang/models"
	"github.com/jinzhu/gorm"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	tkn, err := jwt.GetToken(r)

	if err != nil {
		respond.RespondError(w, http.StatusUnauthorized, "you need to be authenticated to do this")
		return false
	}
	if !tkn.Valid {
		respond.RespondError(w, http.StatusUnauthorized, "invalid token")
		return false
	}

	return true
}

func IsAdmin(w http.ResponseWriter, r *http.Request) bool {
	if !IsAuthenticated(w, r) {
		return false
	}
	role, err := jwt.GetRole(r)

	if err != nil {
		respond.RespondError(w, http.StatusUnauthorized, err.Error())
		return false
	}
	if role != models.AdminRole {
		respond.RespondError(w, http.StatusUnauthorized, "you need to be administrator to do this")
		return false
	}

	return true
}

func Signin(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{}
	if err := db.First(&user, models.User{Username: creds.Username}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, "Username doesn't exist")
		return
	}

	if !password.CheckPasswordHash(creds.Password, user.Password) {
		respond.RespondError(w, http.StatusUnauthorized, "Wrong password")
		return
	}

	token, err := jwt.SetToken(user, w)

	if err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	respond.RespondJSON(w, http.StatusCreated, token)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	if IsAuthenticated(w, r) {
		respond.RespondJSON(w, http.StatusFound, "You are authenticated")
	}
}

func CheckAdmin(w http.ResponseWriter, r *http.Request) {
	if IsAdmin(w, r) {
		respond.RespondJSON(w, http.StatusFound, "You are administrator")
	}
}
