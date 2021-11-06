package main

import (
	"antri-in-backend/api"
	"antri-in-backend/config"
	"antri-in-backend/db"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Inisialisasi Logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Inisialisasi DB
	config := config.GetConfig()
	db.Init(config.DatabaseInit.RemoveAllTables, config.DatabaseInit.InitTestAccount)
	// Inisialisasi Server
	e := api.Init()

	// Server Listener
	e.Logger.Fatal((e.Start(":" + config.ServicePort)))
}
