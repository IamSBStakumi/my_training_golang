package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


type MessageJSONBody = struct {
	Message string
}

func (*Server) SayHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, MessageJSONBody{
		Message: "Hello!",
	})
}