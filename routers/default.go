package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	return router
}
