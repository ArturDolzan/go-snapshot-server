package router

import (
	routerhandler "github.com/arturdolzan/go-snapshot-server/routeshandler"

	"github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine) {
	routerhandler.Initialize()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/nodes", routerhandler.DescribeNodes)
		v1.GET("/images", routerhandler.DescribeImages)
	}
}
