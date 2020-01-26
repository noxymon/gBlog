package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"noxymon.web.id/internal/blog/service/homepage"
)

func DisplayHome(c echo.Context) error {
	homepageGenerator := createNewHomePageGenerator("/home/noxymon/geeblog-content/posts", 4, 250)
	var homepageContent = homepage.GenerateHomepage(homepageGenerator)
	return c.Render(http.StatusOK, "template/blog/index.html", echo.Map{
		"contents": homepageContent,
	})
}

func createNewHomePageGenerator(pathToContent string, numOfPost int, wordToTruncate int) homepage.HomepageGenerator {
	generator := homepage.HomepageGenerator{
		PathToContent:       pathToContent,
		NumPost:             numOfPost,
		NumOfWordToTruncate: wordToTruncate,
	}
	return generator
}
