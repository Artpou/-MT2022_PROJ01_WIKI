type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatorID uint
	ArticleID uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}