package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/h8gi/ocr/models"
	"github.com/labstack/echo"
)

func ShowAllFiles(c echo.Context) error {
	var files []*models.File
	DB.Find(&files)
	return c.JSON(http.StatusOK, files)
}

// Upload file
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
	path := "uploads/" + file.Filename
	f := new(models.File)
	// if record exists
	if !DB.Find(f, "name = ?", file.Filename).RecordNotFound() {
		return c.String(http.StatusConflict, "file already exists.")
	}

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	size, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	f = &models.File{
		Name: file.Filename,
		Path: path,
		Type: file.Header.Get("Content-Type"),
		Size: size,
	}
	if err := DB.Create(f).Error; err != nil {
		return err
	}
	return c.String(http.StatusOK, f.Text)
}

// get file contents (GET /api/files/:name)
func GetFile(c echo.Context) error {
	name := c.Param("name")
	f := new(models.File)
	if DB.First(f, "name = ?", name).RecordNotFound() {
		return c.String(http.StatusNotFound, "file not found")
	}
	return c.File(f.Path)
}

// update file contents (PUT /api/files/:name)
func UpdateFile(c echo.Context) error {
	return c.String(http.StatusServiceUnavailable, "todo: update file")
}

func DeleteFile(c echo.Context) error {
	return c.String(http.StatusServiceUnavailable, "todo: delete file")
}

// get file informations (GET /api/files/:name/info)
func GetFileInfo(c echo.Context) error {
	// name := c.Param("name")
	return c.String(http.StatusServiceUnavailable, "todo: return file info.")
}

// update file information (PUT /api/files/:name/info)
func UpdateFileInfo(c echo.Context) error {
	return c.String(http.StatusServiceUnavailable, "todo: update file info.")
}
