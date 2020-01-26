package post

import (
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"noxymon.web.id/internal/blog/model"
	"time"
)

func GetContent(path string, id string) *model.PostModel {
	var content = getParsedMarkdownForSinglePost(path + "/" + id + ".md")

	var postModel = new(model.PostModel)
	postModel.Init(id, time.Now())
	postModel.SetContent(string(content))

	return postModel
}

func getParsedMarkdownForSinglePost(path string) []byte {
	singlePostContentByte, _ := ioutil.ReadFile(path)
	singlePostContentParsed := blackfriday.Run(singlePostContentByte)
	return singlePostContentParsed
}
