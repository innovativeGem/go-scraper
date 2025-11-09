package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.booking.com/dealspage.html"

	// Fetch the webpage
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL: ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML: ", err)
	}

	// Extract and print page title
	title := doc.Find("h1").Text()
	fmt.Printf("Page Title: %s\n\n", title)

	// Extract all links
	fmt.Println("Links found:")
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Find("a").Attr("href")

		if exists && strings.HasPrefix(href, "https") {
			linkText := s.Text()
			fmt.Printf("%d: %s -> %s\n", i+1, linkText, href)
		}
	})
}