package controllers

import (
	"database/sql"
	"expensetrackerapi/internal/services"
	"expensetrackerapi/pkg/models"
	"expensetrackerapi/pkg/utils"
	"net/http"
	"strconv"

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

func (c *ExpenseController) GetExpense(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("error converting paramater to number: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid exepnse ID")
		return
	}

	expense, err := c.Service.GetExpenseByID(id)
	if err != nil {
		log.Errorf("error getting blog: %v", err)
		if err == sql.ErrNoRows {
			utils.RespondWithError(ctx, http.StatusNotFound, "Expense not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve expense")
		}
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, expense)
}
