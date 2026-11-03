package main

import (
	"github.com/SAB1825/go_shop/internal/config"
	"github.com/SAB1825/go_shop/internal/database"
	"github.com/SAB1825/go_shop/internal/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.New()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load the config")
	}

	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	mainDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get connection to the database")
	}
	defer mainDB.Close()

	gin.SetMode(cfg.Server.GinMode)

	log.Info().Msg("starting server")
}
