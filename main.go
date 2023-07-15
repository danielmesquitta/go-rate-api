package main

import (
	"github.com/danielmesquitta/go-rate-api/config"
	"github.com/danielmesquitta/go-rate-api/controller"
	"github.com/danielmesquitta/go-rate-api/router"
)

func init() {
	config.Init()
	controller.Init()
}

func main() {
	router.Init()
}
