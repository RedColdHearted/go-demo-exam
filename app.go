package main

import (
	"net/http"

	"github.com/RedColdHearted/go-demo-exam/controllers"
	"github.com/RedColdHearted/go-demo-exam/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	v1 := router.Group("/api/v1")
	v1.GET("reservation", controllers.GetReservations)
	v1.GET("reservation/:id", controllers.GetReservation)
	v1.POST("reservation", controllers.CreateReservation)
	v1.PUT("reservation/:id", controllers.UpdateReservation)
	v1.DELETE("reservation/:id", controllers.DeleteReservation)

	models.ConnectDatabase()

	router.Run()
}
