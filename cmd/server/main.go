package main

import (
	"context"
	"database/sql"
	"errors"
	"expensetrackerapi/config"
	"expensetrackerapi/internal/db"
	"expensetrackerapi/internal/middleware"
	"expensetrackerapi/internal/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize Logrus
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("Could not load config: %v", err)
	}
	logger.Info("Loaded configuration file")

	// Initialize database connection
	database, err := db.InitDB(cfg.DatabaseURL)
	if err != nil {
		logger.Fatalf("Could not connect to the database: %v", err)
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Errorf("Error closing database connection: %v", err)
		}
	}(database)
	logger.Info("Database connection initialized")

	// Set Gin mode
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.Default()

	// Apply global middlewares
	router.Use(middleware.GinLogrus(logger), gin.Recovery())

	// Setup routes
	routes.SetupRoutes(router, database)

	// Start the server with graceful shutdown
	port := cfg.Port
	if port == "" {
		port = "8080" // Default to port 8080
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Start server in a separate goroutine
	go func() {
		logger.Infof("Server running on port %s", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("Could not start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shut down: %v", err)
	}
	logger.Info("Server exited properly")
}
