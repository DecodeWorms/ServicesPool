package storage

import (
	"context"
	"servicesPool/models"
)

type UserServices interface {
	AutoMigrateTable(ctx context.Context) error
	AutoMigrateAddressTable(ctx context.Context) error
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
