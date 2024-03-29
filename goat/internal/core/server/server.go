package server

import (
	"github.com/gin-gonic/gin"

	"goat/config"
	"goat/internal/core/logger"
	"goat/internal/controller"
)

func Run() {
	cf := config.GetConfig()
	logger.SetOutputLogToFile()
	r := router()
	r.Run(":" + cf.AppPort)
}

func router() *gin.Engine {
	r := gin.Default()
	
	//TEMPLATE
	r.LoadHTMLGlob("web/template/*.html")

	//STATIC
	r.Static("/css", "web/static/css")
	r.Static("/js", "web/static/js")

	controller.SetRouter(r)

	return r
}
