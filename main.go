package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/extism/go-pdk"
)

type Scraper struct {
	Html     string `json:"html"`
	Selector string `json:"selector"`
}

type RewriteRule struct {
	Selector    string `json:"selector"`
	HTMLContent string `json:"html_content"`
}

type HTMLRewriterInput struct {
	Html  string        `json:"html"`
	Rules []RewriteRule `json:"rules"`
}

//go:export scraper
func scraper() int32 {
	input := Scraper{}
	err := pdk.InputJSON(&input)
	if err != nil {
		pdk.SetError(err)
	}

	if input.Html == "" || input.Selector == "" {
		pdk.SetError(fmt.Errorf("html and selector are required"))
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(input.Html))
	if err != nil {
		pdk.SetError(err)
	}

	doc.Find(input.Selector).Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		pdk.OutputString(fmt.Sprintf("%s", text))
	})

	return 0
}

//go:export htmlrewrite
func htmlrewrite() int32 {
	input := HTMLRewriterInput{}
	err := pdk.InputJSON(&input)
	if err != nil {
		pdk.SetError(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(input.Html))
	if err != nil {
		pdk.SetError(err)
	}

	for _, rule := range input.Rules {
		doc.Find(rule.Selector).Each(func(_ int, s *goquery.Selection) {
			s.SetHtml(rule.HTMLContent)
		})
	}

	html, err := doc.Html()
	if err != nil {
		pdk.SetError(err)
	}
	pdk.OutputString(html)
	return 0
}

func main() {}
