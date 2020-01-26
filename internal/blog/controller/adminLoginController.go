package controller

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"noxymon.web.id/internal/blog/service/auth"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type adminLoginFailedResponse struct {
	Message string `json:"message"`
}

func DisplayAdminLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "template/admin/login.html", echo.Map{})
}

func AuthAdminLogin(c echo.Context) (err error) {
	loggedInUser, done := getLoggedInUserFromRequest(c)

	if done != nil {
		return
	}

	userLoginParam := composeUserLoginParam(loggedInUser)

	if auth.VerifyUser(userLoginParam) {
		saveSession(c)
		return c.JSON(http.StatusAccepted, nil)
	} else {
		return c.JSON(http.StatusForbidden, adminLoginFailedResponse{"no match user & password !"})
	}
}

func saveSession(c echo.Context) {
	sess, _ := session.Get("session", c)
	sess.Values["loggedInUser"] = true
	sess.Save(c.Request(), c.Response())
}

func getLoggedInUserFromRequest(c echo.Context) (u *User, err error) {
	loggedInUser := new(User)
	if err = c.Bind(loggedInUser); err != nil {
		return
	}
	return loggedInUser, nil
}

func composeUserLoginParam(loggedInUser *User) auth.UserLoginParam {
	userLoginParam := auth.UserLoginParam{
		loggedInUser.Username,
		loggedInUser.Password,
	}
	return userLoginParam
}
