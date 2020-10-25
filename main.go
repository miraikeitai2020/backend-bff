package main

import(
	"github.com/miraikeitai2020/backend-bff/pkg/config"
	"github.com/miraikeitai2020/backend-bff/pkg/server"
	"github.com/miraikeitai2020/backend-bff/pkg/server/controller"
)

func main() {
	addr := config.GetRouterAddr()

	ctrl := controller.Init()
	if err := server.Router(ctrl).Run(addr); err != nil {
		panic(err)
	}
}