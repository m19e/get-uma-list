package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	url := "https://ja.wikipedia.org/wiki/%E6%9C%AA%E7%A2%BA%E8%AA%8D%E7%94%9F%E7%89%A9%E4%B8%80%E8%A6%A7"

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	selector := "div.div-col.columns.column-count.column-count-3 > ul > li"

	// Find the review UMA names
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		// For each UMA found, get the name
		name := s.Text()

		name = removeBrackets(name, "（")
		name = removeBrackets(name, "(")
		name = strings.TrimRight(name, " ")
		fmt.Printf("%s\n", name)
	})

	// Why do I need this process?
	doc.Find("#mw-content-text > div > ul:nth-child(14) > li").Each(func(i int, s *goquery.Selection) {
		// For each UMA found, get the name
		name := s.Text()

		name = removeBrackets(name, "（")
		name = removeBrackets(name, "(")
		name = strings.TrimRight(name, " ")
		fmt.Printf("%s\n", name)
	})

	return nil
}

func removeAddendum(s string, a string) string {
	return strings.NewReplacer(
		a, "",
	).Replace(s)
}

func removeBrackets(text string, sep string) string {
	return strings.Split(text, sep)[0]
}
