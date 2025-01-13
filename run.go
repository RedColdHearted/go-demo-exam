package main

import (
	"log"

	"github.com/RedColdHearted/go-demo-exam/models"
	"github.com/RedColdHearted/go-demo-exam/routers"
	"github.com/RedColdHearted/go-demo-exam/routers/api"
	"github.com/RedColdHearted/go-demo-exam/app"
)

var App *app.Application

func main() {
    db, err := models.ConnectDatabase()
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
	r := routers.CommonRouter()
    App := &app.Application{
		DB:     db,
		Router: r,
	}
    api.SetupV1Api(App)

    if err := App.Run(); err != nil {
        log.Fatalf("Could not run server: %v", err)
    }
}
