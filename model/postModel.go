package model

import (
	"errors"

	"gorm.io/gorm"
	"www.blog.com/config"
)

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

// By using association belongs to
// combining the User model and post to extract the data from the database based on user
// Association
type PostUser struct {
	Post `gorm:"embedded"`
	//references is stating that the id in user model is user_id in post model
	User UserData `gorm:"references:id;foreignKey:user_id"`
}

func FindPostById(post_Id uint) (*PostUser, error) {
	var post_user *PostUser = new(PostUser)

	db := config.GoConnect().Debug()
	//Preload loads the data according to association
	if result := db.Where("id=? AND active =?", post_Id, 1).Preload("User").Take(&post_user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return post_user, nil

}

// to get the data of just posts based on post id we use
// func FindPostDataById(post_id int) (*Post, error) {
// 	var post_user *Post = new(Post)

// 	db := config.GoConnect().Debug()
// 	if result := db.Where("id=? AND active=?", post_id, 1).Take(&post_user); result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

// 			return nil, nil
// 		}
// 		return nil, result.Error
// 	}
// 	return post_user, nil

// }
