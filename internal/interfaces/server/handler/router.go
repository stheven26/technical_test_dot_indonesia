package handler

import (
	"technical-test/internal/interfaces/container"

	"github.com/labstack/echo/v4"
)

func SetupRouter(server *echo.Echo, container container.Container) {
	handler := SetupHandler(container).Validate()
	student := server.Group("/student")
	{
		student.POST("", handler.studentHandler.CreateStudent)
		student.GET("", handler.studentHandler.GetAll)
		student.GET("/:id", handler.studentHandler.GetOneById)
		student.PUT("/:id", handler.studentHandler.UpdateById)
		student.PATCH("/:id", handler.studentHandler.PatchById)
		student.DELETE("/:id", handler.studentHandler.DeleteById)
	}
	class := server.Group("/class")
	{
		class.POST("", handler.classHandler.CreateClass)
		class.GET("", handler.classHandler.GetAll)
		class.GET("/:id", handler.classHandler.GetOneById)
		class.DELETE("/:id", handler.classHandler.DeleteById)
	}
}
