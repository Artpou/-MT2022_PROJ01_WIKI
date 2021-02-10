# WIKI_GOLANG API DOCUMENTATION

## Open Endpoints

* [Login] : `POST /api/login/`
* [Articles]: `GET /api/articles/`
* [Article]: `GET /api/article/:id_article


## Endpoints that require Authentication

### Article related

* [Add Article]: `POST /api/article/`
* [Delete Article]: `DELETE /api/article/`
* [Modify Article]: `PUT /api/article/`

* [Add Comment]: `POST /api/comments/`
* [Delete Comment]: `DELETE /api/comments/`
* [Modify Comment]: `PUT /api/comments/`

### Current User related

* [Show Info]: `POST /api/user/`
* [Delete self]: `DELETE /api/user/`
* [Update info]: `PUT /api/user/`
