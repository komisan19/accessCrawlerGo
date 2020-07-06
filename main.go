package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://github.com/tokyo-metropolitan-gov/covid19/blob/development/FORKED_SITES.md"

func checkUrl(URL string) {
	doc, _ := goquery.NewDocument(URL)
	doc.Find("table").Children().Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
          fmt.Println(url)
	})
}

func main() {
	checkUrl(URL)
}
