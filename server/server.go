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
			c.JSON(http.StatusInternalServerError, "err migrating table")
			return
		}
		c.JSON(200, "table created successfully..")
	}
}
