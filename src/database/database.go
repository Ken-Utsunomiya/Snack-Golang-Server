package database

import (
	"Snack-Golang-Server/src/utils"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	dbHost := utils.GetEnvVariable("DB_HOST")
	dbUser := utils.GetEnvVariable("DB_USERID")
	dbName := utils.GetEnvVariable("DB_NAME")
	dbPassword := utils.GetEnvVariable("DB_PASSWORD")
	dbPort := utils.GetEnvVariable("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=", dbHost, dbUser, dbPassword, dbName, dbPort)
	sqlDB, sqlErr := sql.Open("postgres", dsn)
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

	return gormDB
}
