package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Name     string
	Password string
	Mobile   string
}

// create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// get user by id
func GetUser(db *gorm.DB, user *User, email string) (err error) {
	err = db.Where("email= ?", email).First(user).Error
	if err != nil {
		return err
	}
	return nil
}
