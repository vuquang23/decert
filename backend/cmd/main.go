package main

import (
	"decert/internal/pkg/builder"
	"decert/internal/pkg/components"
	"decert/internal/pkg/config"
)

func main() {
	config.Load("")
	components.Init()
	server, _ := builder.NewApiServer()
	server.Run(config.Instance().Http.BindAddress)
}
