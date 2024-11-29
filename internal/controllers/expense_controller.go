package controllers

import (
	"database/sql"
	"errors"
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

	log.Infof("received data to add expense: %+v", expense)

	if err := utils.ValidateExpense(&expense); err != nil {
		log.Errorf("error validating expense: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload: "+error.Error(err))
		return
	}

	err := c.Service.AddExpense(&expense)
	if err != nil {
		log.Errorf("error adding expense: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to add expense")
	}

	log.Infof("expense added for ID %d", expense.ID)

	utils.RespondWithJSON(ctx, http.StatusCreated, expense)
}

func (c *ExpenseController) GetExpense(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("error converting paramater to number: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid expense ID")
		return
	}

	log.Infof("received ID %d to retrieve expense", id)

	expense, err := c.Service.GetExpenseByID(id)
	if err != nil {
		log.Errorf("error getting blog: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(ctx, http.StatusNotFound, "Expense not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve expense")
		}
		return
	}

	log.Infof("retrieved expense for ID %d", id)

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

	log.Info("received data to update expense: %+v", expense)

	if err := utils.ValidateExpense(&expense); err != nil {
		log.Errorf("error validating expense: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	expense.ID = id
	if err := c.Service.UpdateExpense(&expense); err != nil {
		log.Errorf("error updating expense: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
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

	log.Infof("expense updated for ID %d", expense.ID)

	utils.RespondWithJSON(ctx, http.StatusOK, updatedExpense)
}

func (c *ExpenseController) DeleteExpense(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("error converting paramater to integer: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid expense ID")
		return
	}

	log.Infof("received ID %d to delete expense", id)

	if err := c.Service.DeleteExpense(id); err != nil {
		log.Errorf("error deleting blog: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(ctx, http.StatusNotFound, "Expense not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to delete blog")
		}
		return
	}

	log.Infof("deleted expense ID %d", id)

	ctx.Status(http.StatusNoContent)
}
