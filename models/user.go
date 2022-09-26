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

// get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
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

// update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

// delete user
func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
