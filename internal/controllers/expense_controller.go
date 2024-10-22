package controllers

import (
	"expensetrackerapi/internal/services"
	"expensetrackerapi/pkg/models"
	"expensetrackerapi/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ExpenseController struct {
	Service services.ExpenseService
}

func NewExpenseController(service services.ExpenseService) *ExpenseController {
	return &ExpenseController{service}
}

func (c *ExpenseController) AddExpense(ctx *gin.Context) {
	var expense models.Expense
	if err := ctx.ShouldBindJSON(&expense); err != nil {
		log.Errorf("error binding JSON: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
	}

	if err := utils.ValidateExpense(&expense); err != nil {
		log.Errorf("error validating expense: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to add expense")
		return
	}

	utils.RespondWithJSON(ctx, http.StatusCreated, expense)
}
