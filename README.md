# WIKI_GOLANG API DOCUMENTATION

## Open Endpoints

* [Login](docs/login.md) : `POST /api/login/`
* [Show Articles](docs/article_show.md) : `GET /api/articles/`
* [Show Article](docs/article_show.md) : `GET /api/article/:id_article`

## Endpoints that require Authentication

### Article related

* [Add Article](docs/article_add.md) : `POST /api/article/`
* [Delete Article](docs/article_delete.md) : `DELETE /api/article/`
* [Update Article](docs/article_update.md) : `PUT /api/article/`

* [Add Comment](docs/comment_add.md) : `POST /api/comments/`
* [Delete Comment](docs/comment_delete.md) : `DELETE /api/comments/`
* [Update Comment](docs/comment_update.md) : `PUT /api/comments/`

### Current User related

* [Show Info](docs/user_show.md) : `POST /api/user/`
* [Delete self](docs/user_delete.md) : `DELETE /api/user/`
* [Update info](docs/user_update.md) : `PUT /api/user/`
