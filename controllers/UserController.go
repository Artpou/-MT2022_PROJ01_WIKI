package controllers

var db;


func addUser(username string, password string)  {
	user = User.create(username,password);
	db.Create(user);
	View.renderUser(user);
}
