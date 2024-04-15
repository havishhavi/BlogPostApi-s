package model

import "www.blog.com/config"

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

// the table name is users when we declare in table database
func (User) Tablename() string {
	return "users"
}

func FindUserByEmail(Email string) (int64, error) {
	var user User
	var count int64
	// gives you where the error is occured  while creating the table and executing it
	//db := config.GoConnect().Debug()
	db := config.GoConnect()
	if result := db.Model(&user).Where("email = ? AND active = ?", Email, 1).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

//log labels : info log , warning log and error log
// loga are used when we are in server and we dont know whats errors are occured so the log file is important
