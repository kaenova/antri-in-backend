package config

import (
	"antri-in-backend/utils/errlogger"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ServicePort  string
	DatabaseInit DatabaseInitialization
	Database     DatabaseConfig
	Secret       string
}

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
	SSLMode  string `env:"DATABASE_SSLMODE, required"`
}

type DatabaseInitialization struct {
	RemoveAllTables bool `env:"REMOVE_ALL_TABLES, default=false"`
	InitTestAccount bool `env:"INIT_TEST_ACCOUNT, default=false"`
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error reading .env files, continuing without dotenv")
	} else {
		log.Info().Msg("ENV Loaded from .env")
	}
	initAccount, err := strconv.ParseBool(os.Getenv("INIT_TEST_ACCOUNT"))
	errlogger.ErrFatalPanic(err)
	removeTable, err := strconv.ParseBool(os.Getenv("REMOVE_ALL_TABLES"))
	errlogger.ErrFatalPanic(err)
	databaseInit := DatabaseInitialization{
		RemoveAllTables: removeTable,
		InitTestAccount: initAccount,
	}

	return Config{
		ServicePort:  os.Getenv("PORT"),
		DatabaseInit: databaseInit,
		Database: DatabaseConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			Username: os.Getenv("DATABASE_USERNAME"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
			SSLMode:  os.Getenv("DATABASE_SSLMODE"),
		},
		Secret: os.Getenv("SECRET"),
	}
}
