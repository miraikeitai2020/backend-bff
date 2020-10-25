package server

import(
	"github.com/gin-gonic/gin"

	"github.com/miraikeitai2020/backend-bff/pkg/server/controller"
)

func Router(ctrl controller.Controller) *gin.Engine {
	router := gin.Default()
	router.GET("/", ctrl.PlaygroundHandler)
	router.POST("/query", ctrl.QueryHandler)
	return router
}
