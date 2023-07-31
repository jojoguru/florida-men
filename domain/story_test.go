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
