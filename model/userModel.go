package model

import (
	"errors"

	"gorm.io/gorm"
	"www.blog.com/config"
)

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

// new user model to extract the details when associated with post model
type UserData struct {
	ID    uint
	Name  string `gorm:"type:varchar(250)"`
	Email string `gorm:"type:varchar(250);unique;not null"`
	//Password string `gorm:"type:varchar(250); not null"`
	//JwtToken string
	Mobile int `gorm:"not null"`
	Active int `gorm:"type:tinyint(10);default:1"`
	Date   `gorm:"embedded"`
}

// the table name is users when we declare in table database
func (User) Tablename() string {
	return "users"
}

// to let the model know that userdata is also a table from users
func (UserData) TableName() string {
	return "users"
}

//one to many relationship

type UserPost struct {
	UserData `gorm:"embedded"`
	Post     []Post `gorm:"reference:id;foreignKey:user_id"`
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

func FindUserDataByEmail(Email string) (*User, error) {
	var user *User = new(User)

	//connect db,run query,check errors, return
	db := config.GoConnect()
	//Take returns record found by db
	if result := db.Where("email= ? AND active = ? ", Email, 1).Take(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func FindUserByID(id uint) (*User, error) {
	var user *User = new(User)

	db := config.GoConnect()
	if result := db.Where("id = ? AND active = ?", id, 1).Take(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func UpdateToken(UserId uint, JwtToken string) (bool, error) {
	db := config.GoConnect()
	data, err := FindUserByID(UserId)
	if err != nil {
		return false, err
	}
	if data != nil {
		data.JwtToken = JwtToken
		// if result := db.Save(&user_data); result.Error != nil {
		// 	return false, result.Error
		// }
		//fmt.Println(user_data)
		if result := db.Model(&User{}).Where("id=?", UserId).Updates(&data); result.Error != nil {
			return false, result.Error
		} else if result.RowsAffected == 0 {
			err := errors.New("record not updated")
			return false, err
		}
		return true, nil
	} else {
		return false, errors.New("user not found")
	}

}

//log labels : info log , warning log and error log
// loga are used when we are in server and we dont know whats errors are occured so the log file is important

func FindAllPostsByUserId(user_id uint) (*UserPost, error) {
	var user_post *UserPost = new(UserPost)
	db := config.GoConnect().Debug()
	if result := db.Where("id=? AND active = ?", user_id, 1).Preload("Post").Take(&user_post); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user_post, nil
}
