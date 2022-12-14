package server

import (
	"technical-test/internal/interfaces/container"
	"technical-test/internal/interfaces/server/handler"
)

func StartService(cont *container.Container) {
	handler.StartHttpService(cont)
}
