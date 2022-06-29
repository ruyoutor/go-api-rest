package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnect() {
	strConnetion := "host=localhost user=postgres password=changeme dbname=alura_go_rest port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConnetion))

	if err != nil {
		log.Panic("Error connecting from database")
	}
}
