package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func Go() {
	url := "https://www.v2ex.com"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".item_title").Each(func(i int, s *goquery.Selection) {
		el := s.Find("a")
		url, _ := el.Attr("href")
		title := el.Text()
		fmt.Printf("%s => https://www.v2ex.com%s\n", title, url)
	})
}

func main() {
	Go()
}
