package handler

import (
	"net/http"
	"regexp"
	"strconv"
	"technical-test/internal/interfaces/usecase/class"
	"technical-test/pkg/constants"

	"github.com/labstack/echo/v4"
	"gitlab.com/gobang/bepkg/response"
	pkgerror "gitlab.com/gobang/error"
)

type classHandler struct {
	classService class.ClassService
}

func NewClassHandler() *classHandler {
	return &classHandler{}
}

func (h *classHandler) SetClassService(repo class.ClassService) *classHandler {
	h.classService = repo
	return h
}

func (h *classHandler) Validate() *classHandler {
	if h.classService == nil {
		panic("classService is nil")
	}
	return h
}

func (h *classHandler) CreateClass(c echo.Context) (err error) {
	ctx := c.Request().Context()
	req := class.ClassRequest{}
	regexpAlSpa := regexp.MustCompile(`^[a-zA-Z ]*$`)
	if !regexpAlSpa.MatchString(req.Subject) {
		err = pkgerror.New(response.ErrorInvalidRequest, "Invalid Subject")
		return
	}
	if err = c.Bind(&req); err != nil {
		return
	}
	res, err := h.classService.CreateClass(ctx, req)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}

func (h *classHandler) GetAll(c echo.Context) (err error) {
	ctx := c.Request().Context()
	res, err := h.classService.GetAll(ctx)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}

func (h *classHandler) GetOneById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	_, err = strconv.Atoi(id)
	if err != nil {
		err = pkgerror.New(response.ErrorInvalidRequest, constants.MESSAGE_FAILED)
		return
	}
	res, err := h.classService.GetOneById(ctx, id)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}

func (h *classHandler) DeleteById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	_, err = strconv.Atoi(id)
	if err != nil {
		err = pkgerror.New(response.ErrorInvalidRequest, constants.MESSAGE_FAILED)
		return
	}
	res, err := h.classService.DeleteById(ctx, id)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}
