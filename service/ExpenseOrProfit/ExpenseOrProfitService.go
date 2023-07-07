package service

import (
	database "finance_backend/database/config"
	models "finance_backend/models/ExpenseOrProfit"
	"time"
)

func CheckIfExistId(id string) (models.ExpenseOrProfit, error) {
	var user models.ExpenseOrProfit
	err := database.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func CreateExpenseOrProfit(epInput models.CreateExpenseOrProfitFields) models.ExpenseOrProfit {
	ep := models.ExpenseOrProfit{
		UserId:          epInput.UserId,
		ExpenseOrProfit: epInput.ExpenseOrProfit,
		ValueToRegister: epInput.ValueToRegister,
		CreatedAt:       time.Now(),
	}

	database.DB.Create(&ep)

	return ep
}
