package domain

import "fmt"

type Story struct {
	Title   string
	Content string
	Link    string
	Date    Date
}

func NewStory(title string, content string, link string, date Date) (*Story, error) {
	if len(title) <= 0 {
		return nil, fmt.Errorf("Title must not be empty")
	}
	if len(content) <= 0 {
		return nil, fmt.Errorf("Content must not be empty")
	}
	if len(link) <= 0 {
		return nil, fmt.Errorf("Link must not be empty")
	}

	return &Story{title, content, link, date}, nil
}
