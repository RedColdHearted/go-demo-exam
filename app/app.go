package app

import (
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

type Application struct {
    DB     *gorm.DB
    Router *gin.Engine
}

func NewApplication(db *gorm.DB) *Application {
    router := gin.Default()
    return &Application{
        DB:     db,
        Router: router,
    }
}


func(application *Application) Run () (error){
	return application.Router.Run()
}
