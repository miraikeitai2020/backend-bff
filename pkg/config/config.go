package config

import(
	"os"
)

func GetRouterAddr() (addr string) {
	port := os.Getenv("PORT")
	addr = ":" + port
	if port == "" {
		addr = ":9020"
	}
	return
}
