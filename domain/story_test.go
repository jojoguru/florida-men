package domain_test

import (
	"testing"

	"github.com/jojoguru/florida-men/domain"
)

func TestTitleMustNotBeEmpty(t *testing.T) {
	_, err := domain.NewStory("", "foo", "http://google.com", domain.Date{})

	if err == nil || err.Error() != "Title must not be empty" {
		t.Fatalf("NewStory() should throw error on empty title")
	}
}

func TestContentMustNotBeEmpty(t *testing.T) {
	_, err := domain.NewStory("foo", "", "http://google.com", domain.Date{})

	if err == nil || err.Error() != "Content must not be empty" {
		t.Fatalf("NewStory() should throw error on empty content")
	}
}

func TestLinkMustNotBeEmpty(t *testing.T) {
	_, err := domain.NewStory("foo", "bar", "", domain.Date{})

	if err == nil || err.Error() != "Link must not be empty" {
		t.Fatalf("NewStory() should throw error on empty link")
	}
}

func TestNewStory(t *testing.T) {
	givenTitle := "foo"
	givenContent := "some content"
	givenLink := "http://google.com"
	givenDate := domain.Date{"January", 1}

	story, err := domain.NewStory(givenTitle, givenContent, givenLink, givenDate)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if story == nil {
		t.Fatalf("Story should not be nil")
	}

	if story.Title != givenTitle {
		t.Fatalf("Incorrect title. Expected %s got %s", story.Title, givenTitle)
	}
	if story.Content != givenContent {
		t.Fatalf("Incorrect content. Expected %s got %s", story.Content, givenContent)
	}
	if story.Link != givenLink {
		t.Fatalf("Incorrect link. Expected %s got %s", story.Link, givenLink)
	}
	if story.Date != givenDate {
		t.Fatalf("Incorrect date")
	}
}
