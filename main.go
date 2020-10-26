package main

import(
	"github.com/miraikeitai2020/backend-bff/pkg/server"
	"github.com/miraikeitai2020/backend-bff/pkg/server/controller"
)

func main() {
	ctrl := controller.Init()
	if err := server.Router(ctrl).Run(); err != nil {
		panic(err)
	}
}