package storage

import (
	"context"
	"servicesPool/models"
)

type UserServices interface {
	AutoMigrateTable(ctx context.Context) error
	AutoMigrateAddressTable(ctx context.Context) error
	AutoMigrateCompanyTable(ctx context.Context) error
	AutoMigratePostUserRegistration(ctx context.Context) error
	AutoMigrateCompanyAddress(ctx context.Context) error
}

type User struct {
	Conn
}

func NewUser(c Conn) User {
	u := User{
		Conn: c,
	}
	return u

}

func (u User) AutoMigrateTable(ctx context.Context) error {
	var user models.User
	return u.Conn.Client.AutoMigrate(&user)

}

func (u User) AutoMigrateAddressTable(ctx context.Context) error {
	var add models.Address
	return u.Conn.Client.AutoMigrate(&add)

}

func (u User) AutoMigrateCompanyTable(ctx context.Context) error {
	var cmp models.Company
	return u.Conn.Client.AutoMigrate(&cmp)
}

func (u User) AutoMigratePostUserRegistration(ctx context.Context) error {
	var post models.PostUserRegistration
	return u.Conn.Client.AutoMigrate(post)

}

func (u User) AutoMigrateCompanyAddress(ctx context.Context) error {
	var address models.CompanyAddress
	return u.Conn.Client.AutoMigrate(address)
}
