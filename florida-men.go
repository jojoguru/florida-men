package main

import (
	"time"

	"github.com/jojoguru/florida-men/domain"
	"github.com/jojoguru/florida-men/fetcher"
	"github.com/jojoguru/florida-men/writer"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {
	date, _ := domain.NewDateFromTime(time.Now())

	fetcher := fetcher.FloridamanbirthdayFetcher{}
	story, err := fetcher.Fetch(date)

	if err != nil {
		return err
	}

	writer := writer.CliWriter{}
	err = writer.Write(story)
	if err != nil {
		return err
	}

	return nil
}
