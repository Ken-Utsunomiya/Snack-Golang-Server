package database

import (
	"Snack-Golang-Server/src/utils"
	"database/sql"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	dbUri := utils.GetEnvVariable("DB_URI")
	connection, err := pq.ParseURL(dbUri)
	if err != nil {
		log.Fatalf("Connection error")
	}
	connection += " sslmode=require"

	sqlDB, sqlErr := sql.Open("postgres", connection)
	if sqlErr != nil {
		log.Fatalf("Error opening sql database")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	gormDB, gormErr := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if gormErr != nil {
		log.Fatalf("Error opening gorm database")
	}

	DB = gormDB
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
