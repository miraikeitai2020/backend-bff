package controller

import(
	"log"
	"errors"
	"context"
	"net/http"
	"runtime/debug"
	"github.com/gin-gonic/gin"

	// generate package
	"github.com/miraikeitai2020/backend-bff/pkg/bff"
	"github.com/miraikeitai2020/backend-bff/pkg/server/controller/resolver"

	// gql-gen package
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type Controller struct {
	PlaygroundServer http.HandlerFunc
	QueryServer http.Handler
}

func Init() Controller {
	srv := handler.NewDefaultServer(bff.NewExecutableSchema(bff.Config{Resolvers: &resolver.Resolver{}}))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	return Controller{
		PlaygroundServer: playground.Handler("GraphQL playground", "/query"),
		QueryServer: queryMiddleware(srv),
	}
}

func (ctrl *Controller) PlaygroundHandler(cx *gin.Context) {
	ctrl.PlaygroundServer(cx.Writer, cx.Request)
}

func (ctrl *Controller) QueryHandler(cx *gin.Context) {
	ctrl.QueryServer.ServeHTTP(cx.Writer, cx.Request)
}

func queryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", r.Header.Get("id"))
		ctx = context.WithValue(ctx, "pass", r.Header.Get("pass"))
		ctx = context.WithValue(ctx, "token", r.Header.Get("token"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
