package routes

import (
	"database/sql"
	"expensetrackerapi/internal/controllers"
	"expensetrackerapi/internal/repository"
	"expensetrackerapi/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	expenseRepo := repository.NewExpenseRepository(db)
	expenseService := services.NewExpenseService(expenseRepo)
	expenseController := controllers.NewExpenseController(expenseService)

	credentialRepo := repository.NewCredentialRepository(db)
	credentialService := services.NewCredentialService(credentialRepo)
	credentialController := controllers.NewCredentialController(credentialService)

	expenseAPI := router.Group("/expenseAPI")

	{
		expense := expenseAPI.Group("/expense")
		{
			expense.POST("/", expenseController.AddExpense)
			expense.GET("/", expenseController.LoadAllExpenses)
			expense.GET("/:id", expenseController.GetExpense)
			expense.PUT("/:id", expenseController.UpdateExpense)
			expense.DELETE("/:id", expenseController.DeleteExpense)
		}
	}

	router.POST("/signup", credentialController.CreateCredential)
	router.POST("/login", credentialController.GetCredential)
}
