package handler

import (
	"context"
	"servicesPool/storage"
)

type UserHandler struct {
	User storage.User
}

func NewUserHandler(u storage.User) *UserHandler {
	user := &UserHandler{
		User: u,
	}
	return user

}

func (u UserHandler) AutoMigrateTable(ctx context.Context) error {
	if err := u.User.AutoMigrateTable(ctx); err != nil {
		return err

	}
	return nil
}

func (u UserHandler) AutoMigrateAddressTable(ctx context.Context) error {
	if err := u.User.AutoMigrateAddressTable(ctx); err != nil {
		return err
	}
	return nil
}
