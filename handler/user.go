package handler

import (
	"context"
	"fmt"
	"servicesPool/storage"
	"servicesPool/util"

	"github.com/go-playground/validator/v10"
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
func (u UserHandler) AutoMigrateCompanyTable(ctx context.Context) error {
	if err := u.User.AutoMigrateCompanyTable(ctx); err != nil {
		return err
	}
	return nil
}



func validateStruct(v util.InitValidator, data interface{}) []error {
	structErr := make([]error, 0)
	if err := v.Struct(data); err != nil {
		for _, value := range err.(validator.ValidationErrors) {
			e := fmt.Errorf("field :%s ", value.StructNamespace())
			structErr = append(structErr, e)
		}
		return structErr
	}
	return nil

