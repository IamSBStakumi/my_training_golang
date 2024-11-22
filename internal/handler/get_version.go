package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type I interface {
	GetVersion(ctx echo.Context) error
	Animal(ctx echo.Context) error
}

type Server struct{}

func NewHandler() I {
	return &Server{}
}

type versionJSONBody = struct {
	Version string
}

func (*Server) GetVersion(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, versionJSONBody{
		Version: "0.10",
	})
}
