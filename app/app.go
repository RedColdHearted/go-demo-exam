package app

import (
	"github.com/RedColdHearted/go-demo-exam/routers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB
	R  *gin.Engine
}

func NewApplication(db *gorm.DB) *Application {
	router := routers.CommonRouter()
	return &Application{
		DB: db,
		R:  router,
	}
}

func (application *Application) Run() error {
	return application.R.Run()
}
