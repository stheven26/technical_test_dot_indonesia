package cmd

import (
	"technical-test/internal/interfaces/container"
	"technical-test/internal/interfaces/server"
)

func Run() {
	container := container.New()
	server.StartService(container)
}
