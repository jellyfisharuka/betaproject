package router

import (
	"betaproject/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/beta/text",handlers.BetaTest)
	return r 
}