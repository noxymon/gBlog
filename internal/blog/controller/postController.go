package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"noxymon.web.id/internal/blog/service/post"
)

func DisplayPost(c echo.Context) error {
	postId := c.Param("id")
	postContent := post.GetContent("/home/noxymon/geeblog-content/posts", postId)
	return c.Render(http.StatusOK, "template/blog/post.html", echo.Map{
		"post": postContent,
	})
}
