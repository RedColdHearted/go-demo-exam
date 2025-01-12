package controllers

import (
	"net/http"

	"github.com/RedColdHearted/go-demo-exam/models"
	"github.com/gin-gonic/gin"
)

func GetReservations(c *gin.Context) {
	var reservations []models.Reservation
	query_result := models.DB.Find(&reservations)
	if query_result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"query": reservations})
}


func GetReservation(c *gin.Context) {
	var reservation models.Reservation

	if err := models.DB.Where(
		"id = ?", c.Param("id"),
		).First(&reservation).Error; err != nil{
			c.JSON(
				http.StatusBadRequest,
				gin.H{"message": "reservation not found"},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"object": reservation,
	})
}


func CreateReservation(c *gin.Context) {
	var input models.CreateReservation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reservation := models.Reservation{
		BaseReservation: input.BaseReservation,
	}
	models.DB.Create(&reservation)
	c.JSON(http.StatusCreated, gin.H{
		"message": "reservation created",
		"object": reservation,
	})
}


func UpdateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := models.DB.Where(
		"id = ?", c.Param("id"),
		).First(&reservation).Error; err != nil{
			c.JSON(
				http.StatusBadRequest,
				gin.H{"message": "reservation not found"},
		)
		return
	}
	var input models.UpdateReservation
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&reservation).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"message": "reservation updated",
		"object": reservation,
	})
}


func DeleteReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := models.DB.Where(
		"id = ?", c.Param("id"),
		).First(&reservation).Error; err != nil{
			c.JSON(
				http.StatusBadRequest,
				gin.H{"message": "reservation not found"},
		)
		return
	}
	models.DB.Delete(&reservation)
	c.JSON(http.StatusOK, gin.H{
		"message": "reservation deleted",
		"object": reservation,
	})
}
