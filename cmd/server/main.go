package main

import (
	"expensetrackerapi/config"
	"expensetrackerapi/internal/db"
	"expensetrackerapi/internal/middleware"
	"expensetrackerapi/internal/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Initialize Logrus
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	log.Info("Loaded config file")

	// Initialize database
	database, err := db.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	err = database.Close()
	if err != nil {
		log.Fatalf("Could not close database: %v", err)
	}

	log.Info("Database initialized")

	// Set Gin to release mode if not in development
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.Default()

	// Add Logrus logging middleware
	router.Use(middleware.GinLogrus(logger), gin.Recovery())

	// Add JWT validation middleware
	router.Use(middleware.AuthMiddleware(), gin.Recovery())

	// Setup routes
	routes.SetupRoutes(router, database)

	// Start server
	log.Printf("Server running on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
