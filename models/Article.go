type User struct {
	ID        uint `gorm:"primaryKey"`
	Password  string
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(user User) {
	return db.Create(user)
}