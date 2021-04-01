package views

import (
	"fmt"

	"github.com/Artpou/wiki_golang/models"
)

func DeleteComment(comment models.Comment) string {
	return fmt.Sprintf(`{Comment "%d" has been deleted}`, comment.ID)
}
