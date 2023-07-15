package main

import (
	"github.com/danielmesquitta/go-rate-api/config"
	"github.com/danielmesquitta/go-rate-api/handler"
	"github.com/danielmesquitta/go-rate-api/router"
)

func init() {
	config.Init()
	handler.Init()
}

func main() {
	router.Init()
}
