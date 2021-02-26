package views

import (
	"fmt"

	"github.com/Artpou/wiki_golang/models"
)

func ShowUser(user models.User) string {
	return fmt.Sprintf(`{ID: %d,"Pseudo: %s, Email: %s"}`, user.ID, user.Pseudo, user.Email)
}

func AddUser(user models.User) string {
	return fmt.Sprintf(`{User "%s" has been created}`, user.Pseudo)
}

func UpdateUser(user models.User) string {
	return fmt.Sprintf(`{User "%s" has been deleted}`, user.Pseudo)
}

func DeleteUser(user models.User) string {
	return fmt.Sprintf(`{User "%s" has been deleted}`, user.Pseudo)
}
