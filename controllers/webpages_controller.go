package controllers

import (
	"net/http"

	"github.com/h8gi/ocr/models"
	"github.com/labstack/echo"
)

func WebIndex(c echo.Context) error {
	var files []*models.File
	DB.Find(&files)
	return c.Render(http.StatusOK, "index", files)
}
