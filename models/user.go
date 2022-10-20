package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId    string `gorm:"user_id"`
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
	Email     string `gormL:"email"`
	Password  string `gorm:"password"`
}

type NewUserInfoVm struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type Address struct {
	gorm.Model
	UserId          string `gorm:"user_id" validate:"required"`
	Name            string `gorm:"name" validate:"required"`
	City            string `gorm:"city" validate:"required"`
	LocalGovernment string `gorm:"lga" validate:"required"`
	PostCode        string `gorm:"post_code" validate:"required"`
}

type AddressInfoVm struct {
	Name            string `json:"name" validate:"required"`
	City            string `json:"city" validate:"required"`
	LocalGovernment string `json:"lga" validate:"required"`
	PostCode        string `json:"post_code" validate:"required"`
}

type PostUserRegistration struct {
	gorm.Model
	UserId      string `gorm:"user_id"`
	Gender      string `gorm:"gender"`
	Title       string `gorm:"title"`
	PhoneNumber string `gorm:"phoneNumber"`
}

type PostRegistrationInfoVm struct {
	Gender      string `json:"gender" validate:"required"`
	Title       string `json:"title" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Address     AddressInfoVm
}
type Company struct {
	gorm.Model
	Company_Id     string `gorm:"user_id"`
	Name           string `gorm:"name"`
	Email          string `gorm:"email"`
	PhoneNumber    string `gorm:"phone_number"`
	CACnumber      string `gorm:"cac_number"`
	NumberOfStaffs int    `gorm:"number_of_staffs"`
	Biography      string `gorm:"biography"`
	Category       string `gorm:"category"`
}

type CompanyAddress struct {
	gorm.Model
	CompanyId       string `gorm:"company_id"`
	Name            string `gorm:"name"`
	City            string `gorm:"city"`
	LocalGovernment string `gorm:"lga"`
	PostCode        string `gorm:"post_code"`
}

type CompanyInfoVm struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"email"`
	PhoneNumber    string `json:"phone_number" validate:"required"`
	CACnumber      string `json:"cac_number" validate:"required"`
	NumberOfStaffs int    `json:"number_of_staffs" validate:"required"`
}
