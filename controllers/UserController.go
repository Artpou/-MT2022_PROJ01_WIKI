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

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	db.Find(&users)
	handler.RespondJSON(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !handler.IsAuthenticated(w, r) {
		return
	}
	rawUser := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawUser); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if rawUser.Username == "" {
		handler.RespondError(w, http.StatusBadRequest, "Username is missing")
		return
	}
	if rawUser.Password == "" {
		handler.RespondError(w, http.StatusBadRequest, "Password is missing")
		return
	}
	user := models.NewUser(rawUser.Username, rawUser.Password)
	if err := db.Save(&user).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	user := models.User{}
	if err := db.First(&user, models.User{ID: uid}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
	oldUser := models.User{}
	newUser := models.User{}
	if err := db.First(&oldUser, models.User{ID: uid}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if newUser.Password == "" {
		handler.RespondError(w, http.StatusBadRequest, "Password is missing")
		return
	}

	updatedUser := models.UpdateUser(oldUser, newUser.Password)

	if err := db.Save(&updatedUser).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, updatedUser)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
	user := models.User{}
	if err := db.First(&user, models.User{ID: uid}).Error; err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusNoContent, nil)
}

func GetSelf(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.ShowUser(*user)))*/
}

func UpdateSelf(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.UpdateUser(*user)))*/
}

func DeleteSelf(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.DeleteUser(*user)))*/
}
