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

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAdmin(w, r) {
		return
	}
	users := []models.User{}
	db.Find(&users)
	respond.RespondJSON(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rawUser := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawUser); err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if rawUser.Username == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("Username"))
		return
	}
	if rawUser.Password == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("Password"))
		return
	}
	user := models.NewUser(rawUser.Username, rawUser.Password)
	if err := db.Save(&user).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusCreated, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAdmin(w, r) {
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
	user := models.User{}
	if err := db.First(&user, models.User{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("User"))
		return
	}
	respond.RespondJSON(w, http.StatusOK, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAdmin(w, r) {
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
	oldUser := models.User{}
	newUser := models.User{}
	if err := db.First(&oldUser, models.User{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("User"))
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if newUser.Password == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("Password"))
		return
	}

	updatedUser := models.UpdateUser(oldUser, newUser.Password)

	if err := db.Save(&updatedUser).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusOK, updatedUser)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAdmin(w, r) {
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
	user := models.User{}
	if err := db.First(&user, models.User{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("User"))
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusNoContent, views.DeleteUser())
}

func GetSelf(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAuth(w, r) {
		return
	}
	claims, err := jwt.GetClaims(r)
	if err != nil {
		respond.RespondError(w, http.StatusUnauthorized, err.Error())
		return
	}
	username := claims.Username
	user := models.User{}
	if err := db.First(&user, models.User{Username: username}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusOK, user)
}

func UpdateSelf(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAuth(w, r) {
		return
	}
	claims, err := jwt.GetClaims(r)
	if err != nil {
		respond.RespondError(w, http.StatusUnauthorized, err.Error())
		return
	}
	username := claims.Username
	oldUser := models.User{}
	newUser := models.User{}
	if err := db.First(&oldUser, models.User{Username: username}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if newUser.Password == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldNotFound("Password"))
		return
	}

	updatedUser := models.UpdateUser(oldUser, newUser.Password)

	if err := db.Save(&updatedUser).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusOK, updatedUser)
}

func DeleteSelf(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !CheckAuth(w, r) {
		return
	}
	claims, err := jwt.GetClaims(r)
	if err != nil {
		respond.RespondError(w, http.StatusUnauthorized, err.Error())
		return
	}
	username := claims.Username
	user := models.User{}
	if err := db.First(&user, models.User{Username: username}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusNoContent, views.DeleteUser())
}
