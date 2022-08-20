package main

import (
	"os"
	"servicesPool/config"
	"servicesPool/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var client *storage.Conn

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
}

func main() {
	router := gin.Default()
	router.Run(":1213")

}
