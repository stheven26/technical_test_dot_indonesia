package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"technical-test/internal/interfaces/usecase/student"
	"technical-test/pkg/constants"

	"github.com/labstack/echo/v4"
	"gitlab.com/gobang/bepkg/response"
	pkgerror "gitlab.com/gobang/error"
)

type studentHandler struct {
	studentService student.StudentService
}

func NewStudentHandler() *studentHandler {
	return &studentHandler{}
}

func (h *studentHandler) SetStudentService(repo student.StudentService) *studentHandler {
	h.studentService = repo
	return h
}

func (h *studentHandler) Validate() *studentHandler {
	if h.studentService == nil {
		panic("studentService is nil")
	}
	return h
}

func (h *studentHandler) CreateStudent(c echo.Context) (err error) {
	ctx := c.Request().Context()
	req := student.StudentRequest{}
	regexpAlSpa := regexp.MustCompile(`^[a-zA-Z ]*$`)
	if !regexpAlSpa.MatchString(req.Name) {
		err = pkgerror.New(response.ErrorInvalidRequest, "Invalid Format Name")
		return
	}
	if !regexpAlSpa.MatchString(req.School) {
		err = pkgerror.New(response.ErrorInvalidRequest, "Invalid Format School Name")
		return
	}
	if err = c.Bind(&req); err != nil {
		fmt.Println("school:", req.School)
		return
	}
	res, err := h.studentService.CreateStudent(ctx, req)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}

func (h *studentHandler) GetAll(c echo.Context) (err error) {
	ctx := c.Request().Context()
	res, err := h.studentService.GetAll(ctx)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}

func (h *studentHandler) GetOneById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	_, err = strconv.Atoi(id)
	if err != nil {
		err = pkgerror.New(response.ErrorInvalidRequest, constants.MESSAGE_FAILED)
		return
	}
	res, err := h.studentService.GetOneById(ctx, id)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}

func (h *studentHandler) UpdateById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	_, err = strconv.Atoi(id)
	if err != nil {
		err = pkgerror.New(response.ErrorInvalidRequest, constants.MESSAGE_FAILED)
		return
	}
	req := student.StudentRequest{}
	regexpAlSpa := regexp.MustCompile(`^[a-zA-Z ]*$`)
	if !regexpAlSpa.MatchString(req.Name) {
		err = pkgerror.New(response.ErrorInvalidRequest, "Invalid Format Name")
		return
	}
	if !regexpAlSpa.MatchString(req.School) {
		err = pkgerror.New(response.ErrorInvalidRequest, "Invalid Format School Name")
		return
	}
	if err = c.Bind(&req); err != nil {
		return
	}
	res, err := h.studentService.UpdateById(ctx, id, req)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}

func (h *studentHandler) PatchById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	_, err = strconv.Atoi(id)
	if err != nil {
		err = pkgerror.New(response.ErrorInvalidRequest, constants.MESSAGE_FAILED)
		return
	}
	req := student.PatchStudentNameRequest{}
	regexpAlSpa := regexp.MustCompile(`^[a-zA-Z ]*$`)
	if !regexpAlSpa.MatchString(req.Name) {
		err = pkgerror.New(response.ErrorInvalidRequest, "Invalid Format Name")
		return
	}
	if err = c.Bind(&req); err != nil {
		return
	}
	res, err := h.studentService.PatchById(ctx, id, req)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}

func (h *studentHandler) DeleteById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	_, err = strconv.Atoi(id)
	if err != nil {
		err = pkgerror.New(response.ErrorInvalidRequest, constants.MESSAGE_FAILED)
		return
	}
	res, err := h.studentService.DeleteById(ctx, id)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, res)
}
