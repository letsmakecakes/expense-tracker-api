package routes

import (
	"database/sql"
	"expensetrackerapi/internal/controllers"
	"expensetrackerapi/internal/middleware"
	"expensetrackerapi/internal/repository"
	"expensetrackerapi/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Initialize repositories, services, and controllers
	expenseRepo := repository.NewExpenseRepository(db)
	expenseService := services.NewExpenseService(expenseRepo)
	expenseController := controllers.NewExpenseController(expenseService)

	credentialRepo := repository.NewCredentialRepository(db)
	credentialService := services.NewCredentialService(credentialRepo)
	credentialController := controllers.NewCredentialController(credentialService)

	// Public routes
	router.POST("/signup", credentialController.CreateCredential)
	router.POST("/login", credentialController.GetCredential)

	// Protected routes under /expenseAPI
	expenseAPI := router.Group("/expenseAPI")
	expenseAPI.Use(middleware.AuthMiddleware()) // Apply the AuthMiddleware here
	{
		expense := expenseAPI.Group("/expense")
		{
			expense.POST("/", expenseController.AddExpense)         // Add an expense
			expense.GET("/", expenseController.LoadAllExpenses)     // List all expenses
			expense.GET("/:id", expenseController.GetExpense)       // Get a specific expense by ID
			expense.PUT("/:id", expenseController.UpdateExpense)    // Update an expense
			expense.DELETE("/:id", expenseController.DeleteExpense) // Delete an expense
		}
	}
}
