package storage

import (
	"fmt"
	"log"
	"servicesPool/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	Client *gorm.DB
}

func NewConn(c config.Config, db *gorm.DB) *Conn {
	fmt.Println("Connecting to DB...")
	var err error
	dsn := fmt.Sprintf("host=%s dbname=%s port=%s", c.DatabaseHost, c.DatabaseName, c.DatabasePort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("error connecting to DB..")
	}
	fmt.Println("Connecting to DB is initialized..")
	cli := &Conn{
		Client: db,
	}
	return cli

}
