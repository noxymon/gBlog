package internal

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"noxymon.web.id/internal/blog/controller"
)

type TemplateRenderer struct {
	templates *template.Template
	debug     bool
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("s3cr3t"))))

	e.Renderer = customGoView()

	e.Static("/", "views/static")

	e.GET("/", controller.DisplayHome)
	e.GET("/post/:id", controller.DisplayPost)

	e.GET("/admin", controller.DisplayAdminDashboard)
	e.GET("/admin/login", controller.DisplayAdminLogin)

	e.POST("/admin/login", controller.AuthAdminLogin)

	return e
}

func customGoView() *echoview.ViewEngine {
	return echoview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		Funcs:        GetCustomFunction(),
		DisableCache: false,
		Delims:       goview.Delims{},
	})
}
