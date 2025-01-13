package api

import (
	"github.com/RedColdHearted/go-demo-exam/app"
	"github.com/RedColdHearted/go-demo-exam/controllers"
)


func SetupV1Api(application *app.Application) {
	v1 := application.Router.Group("/api/v1")
	rc := controllers.MakeReservationController(application)

	v1.GET("reservation", rc.GetReservations)
	v1.GET("reservation/:id", rc.GetReservation)
	v1.POST("reservation", rc.CreateReservation)
	v1.PUT("reservation/:id", rc.UpdateReservation)
	v1.DELETE("reservation/:id", rc.DeleteReservation)
}
