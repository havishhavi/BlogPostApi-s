package model

type User struct {
	ID       uint
	Name     string `gorm:"type:varchar(250)"`
	Email    string `gorm:"type:varchar(250);unique;not null"`
	Password string `gorm:"type:varchar(250); not null"`
	JwtToken string
	Mobile   int `gorm:"not null"`
	Active   int `gorm:"type:tinyint(10);default:1"`
	Date     `gorm:"embedded"`
}

// why this func?
func (User) Tablename() string {
	return "users"
}
