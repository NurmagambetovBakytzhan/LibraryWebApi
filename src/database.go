package src

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"os"
)

func DbSetup() (*sqlx.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	info := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	//info := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", "localhost", "postgres", "123", "library_web_api", "5432")

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, err
	}
	return db, nil
}
