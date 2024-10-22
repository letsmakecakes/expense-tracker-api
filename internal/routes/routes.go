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

	api := router.Group("/api")
	{
		expense := api.Group("/expense")
		{
			expense.POST("/", expenseController.AddExpense)
			expense.GET("/", expenseController.LoadAllExpenses)
			expense.GET("/:id", expenseController.GetExpense)
			expense.PUT("/:id", expenseController.UpdateExpense)
			expense.DELETE("/:id", expenseController.DeleteExpense)
		}
	}
}
