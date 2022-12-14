package handler

import (
	"fmt"
	"technical-test/internal/interfaces/container"
	"technical-test/pkg/config"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
)

func StartHttpService(container *container.Container) {
	e := echo.New()
	SetupRouter(e, *container)

	// * Set PORT
	port := config.LoadEnv().GetString("PORT")
	e.Logger.Fatal(e.Start(":" + port))

	// * Start Service
	color.Println(color.Green(fmt.Sprintf("⇨ server started on %s\n", port)))
	gracehttp.Serve(e.Server)
}
