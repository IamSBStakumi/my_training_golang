package handler

import (
	"fmt"
	oapi "train/api/generated"
	"train/internal/animal"

	"net/http"

	"github.com/labstack/echo/v4"
)

type MessageJSONBody = struct {
	Message string
}

func (*Server) Animal(ctx echo.Context) error {
	req := new(oapi.AnimalJSONBody)

	err := ctx.Bind(req)
	if err != nil {
		return fmt.Errorf("リクエストボディが受け取れません")
	}
	fmt.Print(req.HasLegs, req.Name)

	pet := animal.GiveBirth(req.Name, req.HasLegs)

	err = pet.Walk()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, MessageJSONBody{
			Message: "it doesn't have legs",
		})
	}

	return ctx.JSON(http.StatusOK, MessageJSONBody{
		Message: "Hello!",
	})
}
