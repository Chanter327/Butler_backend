package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

var envPath string = ".env"

func SetConfig() (Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return Config{}, err
	}

	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	conf := Config{
		DBUser: dbUser,
		DBPass: dbPass,
		DBHost: dbHost,
		DBPort: dbPort,
		DBName: dbName,
	}

	return conf, nil
}

func ConnectDB() (*gorm.DB, error) {
	conf, err := SetConfig()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", conf.DBHost, conf.DBUser, conf.DBPass, conf.DBName, conf.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
	return db, nil
}