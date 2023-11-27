package main

import (
	"log"

	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/config"
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/di"
)

func main() {

	cfg, confErr := config.LoadConfig()

	if confErr != nil {
		log.Fatal("cannot load config: ", confErr)
	}
	server, Dierror := di.InitializeApi(cfg)

	if Dierror != nil {
		log.Fatal("cannot start server: ", Dierror)
	} else {
		server.Start()
	}
}
