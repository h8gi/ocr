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
		"host=localhost user=yagihiroki dbname=gomi sslmode=disable password=mypassword")
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

	w := e.Group("/web")
	w.GET("", controllers.WebIndex)

	a := e.Group("/api")
	a.Use(middleware.CORS())
	a.GET("/files", controllers.ShowAllFiles)
	a.POST("/files", controllers.PostFile)

	a.GET("/files/:name", controllers.GetFile)
	a.PUT("/files/:name", controllers.UpdateFile)
	a.DELETE("/files/:name", controllers.DeleteFile)

	a.GET("/files/:name/info", controllers.GetFileInfo)
	a.PUT("/files/:name/info", controllers.UpdateFileInfo)

	e.Logger.Fatal(e.Start(":1323"))
}
