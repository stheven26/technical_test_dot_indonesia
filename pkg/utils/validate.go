package utils

import (
	"technical-test/pkg/log"

	"github.com/labstack/echo/v4"
	"gitlab.com/gobang/bepkg/response"
	pkgerror "gitlab.com/gobang/error"
)

func Validate(c echo.Context, s interface{}) (err error) {
	ctx := c.Request().Context()

	if err = c.Bind(s); err != nil {
		log.Error(ctx, "error bind", err.Error())
		err = pkgerror.New(response.ErrorInvalidJson, err.Error())
		return
	}

	log.Info(ctx, "Incoming request", s)

	if err = c.Validate(s); err != nil {
		log.Error(ctx, "error validate", err.Error())
		err = pkgerror.New(response.ErrorInvalidJson, err.Error())
		return
	}

	return
}
