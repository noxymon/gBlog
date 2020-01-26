package model

import "time"

type PostModel struct {
	id          string
	title       string
	createdDate time.Time
	author      string
	content     string
}

func (p *PostModel) Init(id string, createdDate time.Time) {
	p.id = id
	p.createdDate = createdDate
}

func (p *PostModel) Content() string {
	return p.content
}

func (p *PostModel) SetContent(content string) {
	p.content = content
}

func (p *PostModel) Author() string {
	return p.author
}

func (p *PostModel) SetAuthor(author string) {
	p.author = author
}

func (p *PostModel) CreatedDate() time.Time {
	return p.createdDate
}

func (p *PostModel) Title() string {
	return p.title
}

func (p *PostModel) SetTitle(title string) {
	p.title = title
}

func (p *PostModel) Id() string {
	return p.id
}
