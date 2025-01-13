package controllers

import (
	"net/http"

	"github.com/RedColdHearted/go-demo-exam/models"
	"github.com/gin-gonic/gin"
	"github.com/RedColdHearted/go-demo-exam/app"
)

type ReservationController struct {
    App *app.Application
}

func MakeReservationController(
	application *app.Application,
) *ReservationController {
    return &ReservationController{App: application}
}

func (rc *ReservationController) GetReservations(c *gin.Context) {
	var reservations []models.Reservation
	queryResult := rc.App.DB.Find(&reservations)
	if queryResult.Error != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"query": reservations})
}

func (rc *ReservationController) GetReservation(c *gin.Context) {
	var reservation models.Reservation

	if err := rc.App.DB.Where(
		"id = ?", c.Param("id"),
	).First(&reservation).Error; err != nil {
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

func (rc *ReservationController) CreateReservation(c *gin.Context) {
	var input models.CreateReservation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reservation := models.Reservation{
		BaseReservation: input.BaseReservation,
	}
	rc.App.DB.Create(&reservation)
	c.JSON(http.StatusCreated, gin.H{
		"message": "reservation created",
		"object":  reservation,
	})
}

func (rc *ReservationController) UpdateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := rc.App.DB.Where(
		"id = ?", c.Param("id"),
	).First(&reservation).Error; err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "reservation not found"},
		)
		return
	}
	var input models.UpdateReservation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rc.App.DB.Model(&reservation).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"message": "reservation updated",
		"object":  reservation,
	})
}

func (rc *ReservationController) DeleteReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := rc.App.DB.Where(
		"id = ?", c.Param("id"),
	).First(&reservation).Error; err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "reservation not found"},
		)
		return
	}
	rc.App.DB.Delete(&reservation)
	c.JSON(http.StatusOK, gin.H{
		"message": "reservation deleted",
		"object":  reservation,
	})
}
