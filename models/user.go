package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	User_Id   string `gorm:"user_id"`
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
	Email     string `gormL:"email"`
	Password  string `gorm:"password"`
	Age       int    `gorm:"age"`
	Gender    string `gorm:"gender"`
	Title     string `gorm:"title"`
}

type NewUserInfoVm struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type AddressInfoVm struct {
	Name            string `json:"name" validate:"required"`
	City            string `json:"city" validate:"required"`
	LocalGovernment string `json:"lga" validate:"required"`
	PostCode        string `json:"post_code" validate:"required"`
}

type Address struct {
	gorm.Model
	User_Id         string `gorm:"user_id"`
	Name            string `gorm:"name"`
	City            string `gorm:"city"`
	LocalGovernment string `gorm:"lga"`
	PostCode        string `gorm:"post_code"`
}
type Company struct {
	gorm.Model
	User_Id        string `gorm:"user_id"`
	Name           string `gorm:"name"`
	Email          string `gorm:"email"`
	PhoneNumber    string `gorm:"phone_number"`
	CACnumber      string `gorm:"cac_number"`
	NumberOfStaffs int    `gorm:"number_of_staffs"`
	Address        string `gorm:"address"`
	Biography      string `gorm:"biography"`
}

type CompanyInfoVm struct {
	Name           string `gorm:"name" validate:"required"`
	Email          string `gorm:"email" validate:"email"`
	PhoneNumber    string `gorm:"phone_number" validate:"required"`
	CACnumber      string `gorm:"cac_number" validate:"required"`
	NumberOfStaffs int    `gorm:"number_of_staffs" validate:"required"`
	Address        string `gorm:"address" validate:"address"`
}
