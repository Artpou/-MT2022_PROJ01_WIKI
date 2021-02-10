package views

import (
	"encoding/json"
	"fmt"

	"github.com/Artpou/wiki_golang/models"
)

type User struct {
	name string
}

func Main() {
	user := models.User{Pseudo: "Frank"}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}

func Test() {
	fmt.Println("test")
	//affichage JSON de l'user
}
