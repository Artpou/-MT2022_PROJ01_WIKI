# WIKI_GOLANG API DOCUMENTATION

This is an API Documentaion of a Wiki. This application was created in Go for a school project, in this app you can read article and comment this. You can login for update article and moderate comments.

## Contributeurs
 
* Arthur Poulin
* Cédric Pierre-Auguste
* Guillaume Marcel

## Accès au trello

https://trello.com/invite/b/JdlJsrku/e7e93345420fb12d04a35cb04a67a613/wiki-go

## Open Endpoints

* [Login](docs/login.md) : `POST /api/login/`
* [Create User](docs/user_add.md) : `POST /api/users/`
* [Show Article](docs/article_show.md) : `GET /api/articles/:id_article`
* [Show Articles](docs/articles_show.md) : `GET /api/articles/`
* [Show Comment](docs/comment_show.md) : `GET /api/comments/:id_comment`
* [Show Comments](docs/comments_show.md) : `GET /api/comments/:id_article`

## Endpoints that require Authentication

### Article related
* [Add Article](docs/article_add.md) : `POST /api/articles/`
* [Delete Article](docs/article_delete.md) : `DELETE /api/articles/:id_article`
* [Update Article](docs/article_update.md) : `PUT /api/articles/:id_article`

### Comment related
* [Add Comment](docs/comment_add.md) : `POST /api/comments/`
* [Delete Comment](docs/comment_delete.md) : `DELETE /api/comments/:id_comment`
* [Update Comment](docs/comment_update.md) : `PUT /api/comments/:id_comment`

### Current User related

* [Show Info](docs/user_show.md) : `GET /api/users/`
* [Delete self](docs/user_delete.md) : `DELETE /api/users/`
* [Update info](docs/user_update.md) : `PUT /api/users/`
