package config

import(
	"os"
	"log"
	"github.com/kelseyhightower/envconfig"
)

type APIPath struct {
	Auth		string	`envconfig:"AUTH_API"`
	User		string	`envconfig:"USER_API"`
	Record		string	`envconfig:"RECORD_API" default:""`
	Memory		string	`envconfig:"MEMORY_API"`
	Collection	string	`envconfig:"COLLECTION_API"`
	Log			string	`envconfig:"LOG_API"`
	Spot		string	`envconfig:"SPOT_API" default:""`
	Evaluation	string	`envconfig:"EVALUATION_API" default:""`
}

func GetRouterAddr() (addr string) {
	port := os.Getenv("PORT")
	addr = ":" + port
	if port == "" {
		addr = ":9020"
	}
	return
}

func GetAPIPath() APIPath {
	var c APIPath
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err)
	}

	return c
}
