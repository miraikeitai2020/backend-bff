package controller

import(
	"context"
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

func QueryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", r.Header.Get("id"))
		ctx = context.WithValue(ctx, "pass", r.Header.Get("pass"))
		ctx = context.WithValue(ctx, "token", r.Header.Get("token"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
