package controller

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DisplayAdminDashboard(c echo.Context) error {
	loggedInUserSession := getCurrentSession(c)
	if loggedInUserSession != nil {
		return c.Render(http.StatusOK, "template/admin/index.html", echo.Map{})
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
}

func getCurrentSession(c echo.Context) interface{} {
	sess, _ := session.Get("session", c)
	loggedInUserSession := sess.Values["loggedInUser"]
	return loggedInUserSession
}
