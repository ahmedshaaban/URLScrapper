package services

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ScrapperResult ...
type ScrapperResult struct {
	PageTitle     string         `json:"page_title"`
	HeadingsCount map[string]int `json:"headings_count"`
	IntExtLinks   int            `json:"internal_external_links_count"`
	InAccessLinks int            `json:"inaccessible_links_count"`
	LoginForm     bool           `json:"contains_login_form"`
}

// ScrapperService ...
type ScrapperService struct {
}

// NewScrapperService ...
func NewScrapperService() *ScrapperService {
	return &ScrapperService{}
}

// Scrap ...
func (s *ScrapperService) Scrap(r io.Reader) (*ScrapperResult, error) {
	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	scrapperResult := &ScrapperResult{
		HeadingsCount: make(map[string]int, 0),
	}

	// Find and print title
	document.Find("title").Each(func(index int, element *goquery.Selection) {
		titleSrc, err := element.Html()
		if err != nil {
			log.Fatal("Error loading title ", err)
		}
		scrapperResult.PageTitle = titleSrc
	})
	if err != nil {
		return nil, err
	}

	for i := 1; i < 7; i++ {
		scrapperResult.HeadingsCount[fmt.Sprintf("h%d", i)] = document.Find(fmt.Sprintf("h%d", i)).Length()
	}

	scrapperResult.IntExtLinks = document.Find("a").Length()

	document.Find("a").Each(func(i int, s *goquery.Selection) {
		_, ok := s.Attr("href")
		if !ok {
			scrapperResult.InAccessLinks++
		}
	})

	document.Find("form").Each(func(i int, s *goquery.Selection) {
		r, _ := s.Html()
		scrapperResult.LoginForm = strings.Contains(r, "login")
	})

	return scrapperResult, nil
}
