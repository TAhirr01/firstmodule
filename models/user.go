package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Age      uint32
	Email    string `gorm:"unique" gorm:"email"`
	Password string
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() uint32 {
	return u.Age
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age uint32) {
	u.Age = age
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}
