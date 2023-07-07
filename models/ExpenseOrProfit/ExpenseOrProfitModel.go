package models

import "time"

type ExpenseOrProfit struct {
	Id              int       `json:"id"`
	UserId          int       `json:"user_id"`
	ExpenseOrProfit int       `json:"expense_or_profit"`
	ValueToRegister int       `json:"value_to_register"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

var ExpensesOrProfits []ExpenseOrProfit

type CreateExpenseOrProfitFields struct {
	UserId          int `json:"user_id" binding:"required"`
	ExpenseOrProfit int `json:"expense_or_profit" binding:"required"`
	ValueToRegister int `json:"value_to_register" binding:"required"`
}

type UpdateExpenseOrProfitFields struct {
	ExpenseOrProfit int `json:"expense_or_profit"`
	ValueToRegister int `json:"value_to_register"`
}
