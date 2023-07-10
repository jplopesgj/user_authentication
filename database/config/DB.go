package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDataBase() {
	dsn := "host=localhost user=postgres password=root dbname=user_auth port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "user_auth.",
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
}
