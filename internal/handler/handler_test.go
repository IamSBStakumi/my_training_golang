package handler

import (
	"net/http"
	oapi "train/api/generated"

	"testing"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	e := echo.New()
	s := NewHandler()

	oapi.RegisterHandlers(e, s)
	t.Run("/versionにGETリクエストを送るとバージョン情報が返ってくる", func(t *testing.T) {
		response := testutil.NewRequest().Get("/version").GoWithHTTPHandler(t, e)
		assert.Equal(t, http.StatusOK, response.Code())
	})

	t.Run("/animalにGETリクエストを送ると挨拶が返ってくる", func(t *testing.T) {
		response := testutil.NewRequest().Get("/hello").GoWithHTTPHandler(t, e)
		assert.Equal(t, http.StatusOK, response.Code())
	})
}
