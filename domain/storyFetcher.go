package domain

type StoryFetcher interface {
	Fetch(date *Date) (*Story, error)
}
