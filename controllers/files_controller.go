package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func ShowAllFiles(c echo.Context) error {
	return c.Render(http.StatusOK, "allfiles", nil)
}

func PostFile(c echo.Context) error {
	// Read files
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("uploads/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func GetFile(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
