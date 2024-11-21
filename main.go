package main

import (
	oapi "train/generated"
	"train/internal/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	server := handler.Server{}

	oapi.RegisterHandlers(e, &server)
	e.Logger.Fatal(e.Start("localhost:7000"))
}
