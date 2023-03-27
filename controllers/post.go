package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createCustomer(c *gin.Context) {
	var customers []Customer
	db.Table("customer").Find(&customers)

	c.JSON(http.StatusOK, gin.H{
		"status":    http.StatusOK,
		"customers": customers,
	})
}
