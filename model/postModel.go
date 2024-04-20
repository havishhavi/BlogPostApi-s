package model

type Post struct {
	ID     uint
	UserID uint   `gorm:"not null"`
	Title  string `gorm:"type:varchar(250);unique;not null"`
	Post   string `gorm:"not null"`
	Active int    `gorm:"type:tinyint(10);default:1"`
	Date   `gorm:"embedded"`
}

func (Post) TableName() string {
	return "posts"
}
