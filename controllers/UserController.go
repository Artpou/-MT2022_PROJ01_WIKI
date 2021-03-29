package controllers

import (
	"encoding/json"
	"fmt"
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
	rawUser := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawUser); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	user := models.NewUser(rawUser.Username, rawUser.Password)
	if err := db.Save(&user).Error; err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	u, err := strconv.ParseUint(key, 10, 64)
	if err == nil {
		u := uint(u)
		article := db.First(models.Article{}, u)
		fmt.Println("Success : getting Article N.", key)
		fmt.Println(article)
		json.NewEncoder(w).Encode(article)
	} else {
		fmt.Println("Error : ID is not an integer", key)
	}
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
