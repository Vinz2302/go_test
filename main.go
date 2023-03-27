package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Customer struct {
	ID           uint
	Name         string
	Nik          uint
	PhoneNumber  uint
	MembershipID uint
}

func main() {
	//dsn := "host=35.187.248.198 user=postgres password=d3v3l0p8015 dbname=trial_week2_4_vincent port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=35.187.248.198 user=postgres password=d3v3l0p8015 dbname=Tral_Week1_Vincent port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic("failed to get sql.db from gorm.db")
		}
		sqlDB.Close()
	}()

	r := gin.Default()
	r.GET("/customers", func(c *gin.Context) {
		var customers []Customer
		db.Table("customer").Find(&customers)

		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"customers": customers,
		})
	})

	r.POST("/customers", func(c *gin.Context) {
		var customer Customer
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request",
			})
			return
		}
		db.Table("customer").Create(&customer)

		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"message":  "Customer created successfully",
			"customer": customer,
		})
	})

	r.PUT("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")

		var customer Customer
		if err := db.Table("customer").Where("id = ?", id).First(&customer).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Customer not found",
			})
			return
		}
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request",
			})
			return
		}
		db.Table("customer").Save(&customer)

		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"message":  "Customer updated successfully",
			"customer": customer,
		})
	})

	r.DELETE("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")

		var customer Customer
		if err := db.Table("customer").Where("id = ?", id).First(&customer).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Customer not found",
			})
			return
		}

		db.Table("customer").Delete(&customer)

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Customer deleted successfully",
		})
	})

	r.Run(":8080")
}
