package controllers

import (
	"github.com/Artpou/wiki_golang/views"
)

func addUser(username string, password string) {
	user = User.create(username, password)
	db.Create(user)
	views.Test()
}
