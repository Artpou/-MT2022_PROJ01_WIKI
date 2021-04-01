package views

func NeedAuthentication() string {
	return "you need to be authenticated to do this"
}

func NeedAdministrator() string {
	return "you need to be administrator to do this"
}

func InvalidToken() string {
	return "the token is invalid"
}

func Logout() string {
	return "you have been successfully logout"
}

func Authenticated() string {
	return "you are connected"
}

func Administrator() string {
	return "you are administrator"
}

func WrongUsername() string {
	return "this username does not exist"
}

func WrongPassword() string {
	return "you enter a wrong password"
}
