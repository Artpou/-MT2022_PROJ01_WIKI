# WIKI_GOLANG API

This api was created in Go for a school project, in this app you can read article and comment this.
The api use JWT Token to login for create article and moderate comments.

##  [Documentation](docs/documentation.md)

You can access to all the documentation [Here](docs/documentation.md)

## Summary

Libraries used for the API :

* [gorm](https://github.com/jinzhu/gorm)
* [jwt](https://github.com/dgrijalva/jwt-go)
* [mux](https://github.com/gorilla/mux)
* [mysql](https://github.com/jinzhu/gorm/dialects/mysql)
* [bcrypt](https://golang.org/x/crypto/bcrypt)

## Use

Launch project with this command in the repo

```bash
go run main.go
```

## Get Started

* To start the API, use the **go run main.go** command.
* You can execute all requests in non-user-restricted endpoints without a token.
* To create an account, use the [Create User](/docs/user_add.md) route.
* Then, use the [Login request](/docs/login.md) route to get a JWT Token.
* You can then use this token as "Bearer Token" Authentication to have access to user-restricted routes, such as article creation and edition.
* Always use raw JSON data for POST and PUT requests instead of form data.
* Admin access :
```json
{
    "Username" : "admin",
    "Password" : "admin"
}
```

## Contributors
 
* Arthur Poulin
* Cédric Pierre-Auguste
* Guillaume Marcel

## Trello Access

https://trello.com/invite/b/JdlJsrku/e7e93345420fb12d04a35cb04a67a613/wiki-go
