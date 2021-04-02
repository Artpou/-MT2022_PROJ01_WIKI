package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Artpou/wiki_golang/handler/jwt"
	"github.com/Artpou/wiki_golang/handler/password"
	"github.com/Artpou/wiki_golang/handler/respond"
	"github.com/Artpou/wiki_golang/models"
	"github.com/Artpou/wiki_golang/views"
	"github.com/jinzhu/gorm"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	tkn, err := jwt.GetToken(r)
	return err == nil && tkn.Valid
}

func IsAdmin(w http.ResponseWriter, r *http.Request) bool {
	claims, err := jwt.GetClaims(r)
	return IsAuthenticated(w, r) && err == nil && claims.Role == models.AdminRole
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
		respond.RespondError(w, http.StatusNotFound, views.WrongUsername())
		return
	}

	if !password.CheckPasswordHash(creds.Password, user.Password) {
		respond.RespondError(w, http.StatusUnauthorized, views.WrongPassword())
		return
	}

	token, err := jwt.SetToken(user, w)

	if err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	respond.RespondJSON(w, http.StatusCreated, token)
}

func Logout(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAuth(w, r) {
		return
	}
	jwt.DeleteToken(w)
	respond.RespondJSON(w, http.StatusFound, views.Logout())
}

func CheckAuth(w http.ResponseWriter, r *http.Request) bool {
	if !IsAuthenticated(w, r) {
		respond.RespondError(w, http.StatusUnauthorized, views.NeedAdministrator())
		return false
	}
	respond.RespondJSON(w, http.StatusFound, views.Authenticated())
	return true
}

func CheckAdmin(w http.ResponseWriter, r *http.Request) bool {
	if !IsAdmin(w, r) {
		respond.RespondError(w, http.StatusUnauthorized, views.NeedAuthentication())
		return false
	}
	respond.RespondJSON(w, http.StatusFound, views.Administrator())
	return true
}
