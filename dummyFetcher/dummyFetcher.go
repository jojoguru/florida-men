package dummyfetcher

import (
	"github.com/jojoguru/florida-men/domain"
)

func Fetch(date domain.Date) (*domain.Story, error) {
	return domain.NewStory("Dummy Title", "Dummy Content", "http://example.com", date)
}
