package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

  var (
	DB *gorm.DB
	err error
  )

func ConectaDB() {
	
	dsn := "host=localhost user=postgres password=123 dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic("Erro ao conectar com o banco")
	}

	
}