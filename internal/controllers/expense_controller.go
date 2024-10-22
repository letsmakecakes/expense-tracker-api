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

func (c *ExpenseController) LoadAllExpenses(ctx *gin.Context) {
	expenses, err := c.Service.LoadAllExpenses()
	if err != nil {
		log.Errorf("error getting expenses: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve expenses")
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, expenses)
}

func (c *ExpenseController) UpdateExpense(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("error updating expense: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid expense ID")
		return
	}

	var expense models.Expense
	if err := ctx.ShouldBindJSON(&expense); err != nil {
		log.Errorf("error binding JSON: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := utils.ValidateExpense(&expense); err != nil {
		log.Errorf("error validating expense: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	expense.ID = id
	if err := c.Service.UpdateExpense(&expense); err != nil {
		log.Errorf("error updating expense: %v", err)
		if err == sql.ErrNoRows {
			utils.RespondWithError(ctx, http.StatusNotFound, "Expense not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to update expense")
		}
		return
	}

	updatedExpense, err := c.Service.GetExpenseByID(id)
	if err != nil {
		log.Errorf("error getting expense: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve updated expense")
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, updatedExpense)
}
