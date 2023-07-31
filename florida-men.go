package main

import (
	"time"

	cliwriter "github.com/jojoguru/florida-men/CliWriter"
	"github.com/jojoguru/florida-men/domain"
	floridamanbirthdayfetcher "github.com/jojoguru/florida-men/floridamanbirthdayFetcher"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {
	date, _ := domain.NewDateFromTime(time.Now())

	story, err := floridamanbirthdayfetcher.Fetch(date)

	if err != nil {
		return err
	}

	err = cliwriter.Write(story)
	if err != nil {
		return err
	}

	return nil
}
