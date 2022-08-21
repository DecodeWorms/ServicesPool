package server

import (
	"net/http"
	"servicesPool/handler"

	"github.com/gin-gonic/gin"
)

type UserServer struct {
	User handler.UserHandler
}

func NewUserServer(u handler.UserHandler) *UserServer {
	us := &UserServer{
		User: u,
	}
	return us
}

func (u UserServer) AutoMigrateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := u.User.AutoMigrateTable(c); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.String(http.StatusOK, "sucessfull")
	}
}

func (u UserServer) AutoMigrateAddressTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := u.User.AutoMigrateAddressTable(c); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.String(http.StatusOK, "address table created successfully")
	}
}

func (u UserServer) AutoMigrateCompanyTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := u.User.AutoMigrateCompanyTable(c); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.String(http.StatusOK, "successful")

	}
}
