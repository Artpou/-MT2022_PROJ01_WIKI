package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Artpou/wiki_golang/handler"
	"github.com/Artpou/wiki_golang/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var jwtKey = []byte("cle_tres_secrete")

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	handler.RespondJSON(w, http.StatusCreated, "Token Created")
}
