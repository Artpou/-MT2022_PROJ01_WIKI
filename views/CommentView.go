package views

import (
	"fmt"

	"github.com/Artpou/wiki_golang/models"
)

func CommentNotFound() string {
	return "this comment cannot be found"
}

func GetComments() string {
	return fmt.Sprintf(`{all comments}`)
}

func AddComment(comment models.Comment) string {
	return fmt.Sprintf(`{Comment "%d" has been created}`, comment.ID)
}

func UpdateComment(comment models.Comment) string {
	return fmt.Sprintf(`{Comment "%d" has been updated}`, comment.ID)
}

func DeleteComment(comment models.Comment) string {
	return fmt.Sprintf(`{Comment "%d" has been deleted}`, comment.ID)
}
