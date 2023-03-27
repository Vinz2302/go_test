package controllers

import (
	"Vinz2302/go_test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createCustomer(c *gin.Context) {
	var customers []models.Customer
	db.Table("customer").Find(&customers)

	c.JSON(http.StatusOK, gin.H{
		"status":    http.StatusOK,
		"customers": customers,
	})
}
