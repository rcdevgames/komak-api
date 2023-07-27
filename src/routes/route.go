package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/thinktive/gopi/src/middlewares"
)

func Init(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Always ON Guys :)",
		})
	})

	v1Router := router.Group("/v1")
	{
		authRouter := v1Router.Group("/auth")
		{
			authRouter.POST("/login")
			authRouter.Use(middlewares.AuthRefresh()).POST("/refresh-token")
		}

	}
}
