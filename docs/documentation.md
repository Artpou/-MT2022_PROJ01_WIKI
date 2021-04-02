# DOCUMENTATION

## Get Started

You can execute all request in open endpoints without being connected.
You can use the [Login request](/login.md) and put the JWT token
to start creating and editing articles or comments.

## Open Endpoints

### Current User related
* [Login](/login.md) : `POST /api/login/`
* [Create User](/user_add.md) : `POST /api/users/`

### Article related
* [Show Article](/article_show.md) : `GET /api/articles/:id_article`
* [Show Articles](/articles_show.md) : `GET /api/articles/`

### Comment related
* [Show Comment](/comment_show.md) : `GET /api/comments/:id_comment`
* [Show Comments](/comments_show.md) : `GET /api/comments/:id_article`

## Endpoints that require Authentication

This endpoints are only available if you are login to the api

### Article related
* [Add Article](/article_add.md) : `POST /api/articles/`
* [Delete Article (only if you are the owner)](/article_delete.md) : `DELETE /api/articles/:id_article`
* [Update Article (only if you are the owner)](/article_update.md) : `PUT /api/articles/:id_article`

### Comment related
* [Add Comment](/comment_add.md) : `POST /api/comments/`
* [Delete Comment (only if you are the owner)](/comment_delete.md) : `DELETE /api/comments/:id_comment`
* [Update Comment (only if you are the owner)](/comment_update.md) : `PUT /api/comments/:id_comment`

### Current User related
* [Logout](/logout.md) : `POST /api/logout/`
* [Show Info](/self_show.md) : `GET /api/self/`
* [Delete self](/self_delete.md) : `DELETE /api/self/`
* [Update info](/self_update.md) : `PUT /api/self/`

## Endpoints that require Administrator Role

This endpoints are only available if you are administrator

### User related
* [Show Users](/user_show.md) : `GET /api/users/`
* [Delete User](/user_delete.md) : `DELETE /api/users/`
* [Update User](/user_update.md) : `PUT /api/users/`

### Article related
You can delete or update any Article

### Comment related
You can delete or update any Comment
