package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB
var err error

func InitDB() {
	dsn := os.Getenv("POSTGRES_DNS")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NameReplacer:  nil,
		},
	})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Database connected.")
	}
}
