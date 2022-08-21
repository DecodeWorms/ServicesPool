package main

import (
	"os"
	"servicesPool/config"
	"servicesPool/handler"
	"servicesPool/server"
	"servicesPool/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var client *storage.Conn
var userStorage storage.User
var userHandler handler.UserHandler
var userServer server.UserServer

func init() {
	_ = godotenv.Load()
	var db *gorm.DB
	h := os.Getenv("DATABASE_HOST")
	dn := os.Getenv("DATABASE_NAME")
	p := os.Getenv("DATABASE_PORT")

	c := config.Config{
		DatabaseHost: h,
		DatabaseName: dn,
		DatabasePort: p,
	}
	client = storage.NewConn(c, db)
	userStorage = storage.NewUser(*client)
	userHandler = *handler.NewUserHandler(userStorage)
	userServer = *server.NewUserServer(userHandler)
}

func main() {
	router := gin.Default()
	router.POST("/migrate", userServer.AutoMigrateTable())
	router.POST("/migrate2", userServer.AutoMigrateAddressTable())
	router.POST("/migrate3", userServer.AutoMigrateCompanyTable())
	router.Run(":1213")
}
