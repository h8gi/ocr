package main

import (
	"net/http"

	"github.com/h8gi/ocr/controllers"
	"github.com/h8gi/ocr/models"
	"github.com/h8gi/ocr/views"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db, err := gorm.Open("postgres",
		"host=localhost user=yagi dbname=gomi sslmode=disable password=mypassword")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Migration
	db.AutoMigrate(&models.File{})
	controllers.SetDB(db)
	t := views.NewTemplate("./views/*.html")
	e := echo.New()
	// register templates
	e.Renderer = t

	// Middleware
	// remove trailing slash. /hello/ -> /hello
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "assets")
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/api/files", controllers.ShowAllFiles)
	e.POST("/api/files", controllers.PostFile)

	e.GET("/api/files/:name", controllers.GetFile)
	e.PUT("/api/files/:name", controllers.UpdateFile)

	e.GET("/api/files/:name/info", controllers.GetFileInfo)
	e.PUT("/api/files/:name/info", controllers.UpdateFileInfo)

	e.Logger.Fatal(e.Start(":1323"))
}
