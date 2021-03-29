package views

import (
	"fmt"

	"github.com/Artpou/wiki_golang/models"
)

func GetArticles(article models.Article) string {
	return fmt.Sprintf(`{ID: %d}`, article.ID)
}
