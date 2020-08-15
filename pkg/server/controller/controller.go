package controller

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	PlaygroundServer http.HandlerFunc
	QueryServer http.Handler
}

func (ctrl *Controller) PlaygroundHandler(cx *gin.Context) {
	ctrl.PlaygroundServer(cx.Writer, cx.Request)
}

func (ctrl *Controller) QueryHandler(cx *gin.Context) {
	ctrl.QueryServer.ServeHTTP(cx.Writer, cx.Request)
}