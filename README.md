# WIKI_GOLANG API DOCUMENTATION

This is an API Documentaion of a Wiki. This application was created in Go for a school project, in this app you can read article and comment this. You can login for update article and moderate comments.

## Contributors
 
* Arthur Poulin
* Cédric Pierre-Auguste
* Guillaume Marcel

## Trello Access

https://trello.com/invite/b/JdlJsrku/e7e93345420fb12d04a35cb04a67a613/wiki-go

## Open Endpoints

### Current User related
* [Login](docs/login.md) : `POST /api/login/`
* [Create User](docs/user_add.md) : `POST /api/users/`

### Article related
* [Show Article](docs/article_show.md) : `GET /api/articles/:id_article`
* [Show Articles](docs/articles_show.md) : `GET /api/articles/`

### Comment related
* [Show Comment](docs/comment_show.md) : `GET /api/comments/:id_comment`
* [Show Comments](docs/comments_show.md) : `GET /api/comments/:id_article`

## Endpoints that require Authentication

This endpoints are only available if you are login to the api

### Article related
* [Add Article](docs/article_add.md) : `POST /api/articles/`
* [Delete Article (only if you are the owner)](docs/article_delete.md) : `DELETE /api/articles/:id_article`
* [Update Article (only if you are the owner)](docs/article_update.md) : `PUT /api/articles/:id_article`

### Comment related
* [Add Comment](docs/comment_add.md) : `POST /api/comments/`
* [Delete Comment (only if you are the owner)](docs/comment_delete.md) : `DELETE /api/comments/:id_comment`
* [Update Comment (only if you are the owner)](docs/comment_update.md) : `PUT /api/comments/:id_comment`

### Current User related
* [Logout](docs/logout.md) : `POST /api/logout/`
* [Show Info](docs/user_show.md) : `GET /api/self/`
* [Delete self](docs/user_delete.md) : `DELETE /api/self/`
* [Update info](docs/user_update.md) : `PUT /api/self/`

## Endpoints that require Administrator Role

This endpoints are only available if you are administrator

### User related
* [Show Users](docs/user_show.md) : `GET /api/users/`
* [Delete User](docs/user_delete.md) : `DELETE /api/users/`
* [Update User](docs/user_update.md) : `PUT /api/users/`

### Article related
You can delete or update any Article

### Comment related
You can delete or update any Comment
