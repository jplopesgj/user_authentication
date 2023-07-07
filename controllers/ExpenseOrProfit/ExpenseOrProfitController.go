package controllers

import (
	database "finance_backend/database/config"
	models "finance_backend/models/ExpenseOrProfit"
	service "finance_backend/service/ExpenseOrProfit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSpecificExpenseOrProfit(c *gin.Context) {
	id := c.Param("id")
	user, err := service.CheckIfExistId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateExpenseOrProfit(c *gin.Context) {
	var epInput models.CreateExpenseOrProfitFields
	if err := c.ShouldBindJSON(&epInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := service.CreateExpenseOrProfit(epInput)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateExpenseOrProfit(c *gin.Context) {
	id := c.Param("id")
	ep, err := service.CheckIfExistId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var userInput models.UpdateExpenseOrProfitFields
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&ep).Updates(userInput)
	c.JSON(http.StatusOK, gin.H{"data": ep})
}

func DeleteExpenseOrProfit(c *gin.Context) {
	id := c.Param("id")
	ep, err := service.CheckIfExistId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&ep)
	c.JSON(http.StatusOK, gin.H{"data": "deleted"})
}
