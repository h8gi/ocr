package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func ShowAllFiles(c echo.Context) error {
	return c.Render(http.StatusOK, "allfiles", nil)
}
