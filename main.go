package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/extism/go-pdk"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
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
	pdk.InputJSON(&input)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(input.Html))
	if err != nil {
		log.Fatal(err)
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
	pdk.InputJSON(&input)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(input.Html))
	if err != nil {
		log.Fatal(err)
	}

	for _, rule := range input.Rules {
		doc.Find(rule.Selector).Each(func(_ int, s *goquery.Selection) {
			s.SetHtml(rule.HTMLContent)
		})
	}

	html, err := doc.Html()
	if err != nil {
		log.Fatal(err)
	}
	pdk.OutputString(html)
	return 0
}

//go:export md2html
func md2html() int32 {
	input := pdk.InputString()

	var html bytes.Buffer

	gm := goldmark.New(
		goldmark.WithExtensions(
			extension.Linkify,
			extension.Strikethrough,
			extension.Table,
		),
	)
	_ = gm.Convert([]byte(input), &html)

	escapedHTML := strconv.QuoteToASCII(html.String())
	escapedHTML = escapedHTML[1 : len(escapedHTML)-1] // Remove the extra double quotes

	pdk.OutputString(html.String())
	return 0
}

func main() {}
