# WIKI_GOLANG API

This api was created in Go for a school project, in this app you can read article and comment this.
The api use JWT Token to login for create article and moderate comments.

##  [Documentation](docs/documentation.md)

You can access to all the documentation [Here](docs/documentation.md)

## Use

Launch project with this command in the repo

```bash
go run main.go
```

## Get Started

* You can execute all requests in non-user-restricted endpoints without a token.
* To create an account, use the [Create User](/docs/user_add.md) route.
* Then, use the [Login request](/docs/login.md) route to get a JWT Token.
* You can then use this token as "Bearer Token" Authentication to have access to user-restricted routes, such as article creation and edition.
* Always use raw JSON data for POST and PUT requests instead of form data.
* Admin access :
* Username : admin
* Password : admin

## Contributors
 
* Arthur Poulin
* Cédric Pierre-Auguste
* Guillaume Marcel

## Trello Access

https://trello.com/invite/b/JdlJsrku/e7e93345420fb12d04a35cb04a67a613/wiki-go
