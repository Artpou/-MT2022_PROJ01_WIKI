package controllers

import (
	"encoding/json"
	_"fmt"
	_"io/ioutil"
	"net/http"
	_"strconv"

	"github.com/Artpou/wiki_golang/models"
	_"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
  db.Find(&users)
  json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.AddUser(*user)))*/
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.ShowUser(*user)))*/
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.DeleteUser(*user)))*/
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.DeleteUser(*user)))*/
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
