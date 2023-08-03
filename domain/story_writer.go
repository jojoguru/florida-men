package domain

type StoryWriter interface {
	Write(story *Story) error
}
