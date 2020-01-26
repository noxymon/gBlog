package homepage

import (
	"bytes"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"noxymon.web.id/internal/blog/model"
	"os"
	"path/filepath"
	"strings"
)

type HomepageGenerator struct {
	PathToContent       string
	NumPost             int
	NumOfWordToTruncate int
}

func GenerateHomepage(h HomepageGenerator) []*model.PostModel {
	var postsContent []*model.PostModel
	filepath.Walk(h.PathToContent, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			postId := strings.TrimRight(info.Name(), ".md")
			parsedMarkdownContent := getParsedMarkdownFile(postId, path, h.NumOfWordToTruncate)

			var postModel = new(model.PostModel)
			postModel.Init(postId, info.ModTime())
			postModel.SetContent(string(parsedMarkdownContent))

			postsContent = append(postsContent, postModel)
		}
		return nil
	})
	return postsContent
}

func getParsedMarkdownFile(id string, path string, numOfWordToTruncate int) string {
	singlePostContentByte, _ := ioutil.ReadFile(path)
	singlePostContent := getTruncatedPost(singlePostContentByte, numOfWordToTruncate)

	singlePostContentParsed := string(blackfriday.Run(singlePostContent))
	singlePostContentParsed = createReadMore(singlePostContentParsed, id)
	return singlePostContentParsed
}

func createReadMore(contentParsed string, id string) string {
	contentParsed = contentParsed + "<a href=\"./post/" + id + "\">Read more....</a>"
	return contentParsed
}

func getTruncatedPost(content []byte, numOfWordToTruncate int) []byte {
	arrayOfWord := bytes.Split(content, []byte(" "))
	postlength := len(arrayOfWord)
	if postlength < numOfWordToTruncate {
		numOfWordToTruncate = postlength
	}
	return bytes.Join(arrayOfWord[0:numOfWordToTruncate], []byte(" "))
}
