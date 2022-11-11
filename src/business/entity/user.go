package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"-"`
	PhoneNumber string `json:"phone_number"`
}

type CreateUserParam struct {
	Email       string `binding:"required"`
	Password    string `binding:"required"`
	PhoneNumber string `binding:"required" json:"phone_number"`
}

type LoginUserParam struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type SelectUserParam struct {
	Id    int
	Email string
}
