package fetcher

import (
	"github.com/jojoguru/florida-men/domain"
)

type DummyFetcher struct {
}

func (*DummyFetcher) Fetch(date domain.Date) (*domain.Story, error) {
	return domain.NewStory("Dummy Title", "Dummy Content", "http://example.com", date)
}
