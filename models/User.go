type User struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Pseudo    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func create(pseudo string, password string) {
	return User{ 
		Pseudo: pseudo,
		password: password
	}	
}