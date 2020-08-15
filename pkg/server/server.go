package main

import(
	"os"
	"log"
	"errors"
	"context"
	"runtime/debug"
	"github.com/gin-gonic/gin"

	// generate package
	"github.com/SaKu2110/gql-practice/pkg/bff"
	"github.com/SaKu2110/gql-practice/pkg/server/controller"
	"github.com/SaKu2110/gql-practice/pkg/server/resolver"

	// gql-gen package
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "9020"
	}
}

func initializeController() controller.Controller {
	srv := handler.NewDefaultServer(bff.NewExecutableSchema(bff.Config{Resolvers: &resolver.Resolver{}}))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	return controller.Controller{
		PlaygroundServer: playground.Handler("GraphQL playground", "/query"),
		QueryServer: srv,
	}
}

func setupRouter(ctrl controller.Controller) *gin.Engine {
	router := gin.Default()
	router.GET("/", ctrl.PlaygroundHandler)
	router.POST("/query", ctrl.QueryHandler)
	return router
}

func main() {
	ctrl := initializeController()
	router := setupRouter(ctrl)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("failed launch router. err=%s", err)
		os.Exit(1)
	}
}