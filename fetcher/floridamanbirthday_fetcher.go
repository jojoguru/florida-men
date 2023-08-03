package fetcher

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/jojoguru/florida-men/domain"
)

type FloridamanbirthdayFetcher struct{}

func (*FloridamanbirthdayFetcher) Fetch(date domain.Date) (*domain.Story, error) {
	url := buildUrl(date)

	// response, err := fetchData(url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)

		return nil, err
	}

	// Load the HTML document
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	post := document.Find("#the-post").First()
	title := post.Find(".entry-content p b").Text()

	return domain.NewStory(title, "foo", url, date)
}

func buildUrl(date domain.Date) string {
	var baseUrl = "https://floridamanbirthday.org/"

	return fmt.Sprintf("%s/%s-%d", baseUrl, date.Month, date.Day)
}
