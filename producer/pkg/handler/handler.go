package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouters(app *gin.Engine) {
	v1 := app.Group("/v1")
	{
		v1.GET("/ping", Ping)
		v1.POST("/publish/example", SendMessageInQ)
	}

}
