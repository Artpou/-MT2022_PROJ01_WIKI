package views

import (
	"fmt"

	"github.com/Artpou/wiki_golang/models"
)

func ShowUser(user models.User) string {
	return fmt.Sprintf(`{ID: %d,"Username": %s, "Email": %s}`, user.ID, user.Username, user.Email)
}

func AddUser(user models.User) string {
	return fmt.Sprintf(`{User "%s" has been created}`, user.Username)
}

func UpdateUser(user models.User) string {
	return fmt.Sprintf(`{User "%s" has been deleted}`, user.Username)
}

func DeleteUser(user models.User) string {
	return fmt.Sprintf(`{User "%s" has been deleted}`, user.Username)
}
