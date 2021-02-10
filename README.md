# WIKI_GOLANG API DOCUMENTATION

## Open Endpoints

* [Login](docs/login.md) : `POST /api/login/`
* [Create Account](docs/user_add.md) : `POST /api/users/`
* [Show Articles](docs/article_show.md) : `GET /api/articles/`
* [Show Article](docs/article_show.md) : `GET /api/articles/:id_article`

## Endpoints that require Authentication

### Article related

* [Add Article](docs/article_add.md) : `POST /api/articles/`
* [Delete Article](docs/article_delete.md) : `DELETE /api/articles/`
* [Update Article](docs/article_update.md) : `PUT /api/articles/`

* [Add Comment](docs/comment_add.md) : `POST /api/comments/`
* [Delete Comment](docs/comment_delete.md) : `DELETE /api/comments/`
* [Update Comment](docs/comment_update.md) : `PUT /api/comments/`

### Current User related

* [Show Info](docs/user_show.md) : `GET /api/users/`
* [Delete self](docs/user_delete.md) : `DELETE /api/users/`
* [Update info](docs/user_update.md) : `PUT /api/users/`
