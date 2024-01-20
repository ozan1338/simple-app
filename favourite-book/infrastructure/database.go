package infrastructure

import (
	"log"
	"os"

	"favourite-book/domain/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
    dsn := os.Getenv("DB_DSN")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    if err := db.AutoMigrate(&entity.Book{}); err != nil {
        log.Fatal(err)
    }

    return db
}