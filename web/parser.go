package web

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gotk3/gotk3/gdk"
)

type Parser struct {
	document *goquery.Document
}

func NewParser() *Parser {
	parser := new(Parser)
	return parser
}

func (p *Parser) ParseHTML(document string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(document))
	if err != nil {
		log.Fatal(err)
	}
	p.document = doc
}

func (p *Parser) Extract() *[]VpnEntry { // Too hacky stuff, i should rewrite it someday
	in := 0           //yeah its smelly
	StartHit := false //do not try to debug just rewrite it
	EndHit := false
	skipFirst := false
	entries := []VpnEntry{}
	entry := VpnEntry{}
	p.document.Find("div div div").Each(func(i int, s *goquery.Selection) {

		_ = entry
		img := s.Find("img")
		if attr, _ := img.Attr("align"); attr == "left" {
			StartHit = true

		}

		note := s.Find(".note")
		n, _ := note.Html()
		if len(n) > 1 {
			EndHit = true
		}

		if !StartHit || EndHit {
			return
		}

		//title, _ := s.Html()

		if img.Is("img") {
			src, _ := img.Attr("src")
			fmt.Println("img", src)
			if skipFirst {
				entries = append(entries, entry)
			}
			entry = VpnEntry{}

			img, e := DownloadFile(BaseUrl + src)

			if e != nil {
				fmt.Println(e)
			}
			entry.Image, e = gdk.PixbufNewFromBytesOnly(*img)
			if e != nil {
				fmt.Println(e)
			}
			skipFirst = true
		}

		contryList := s.Find(".country_list")
		if contryList.Is(".country_list") {
			fmt.Println("name", contryList.Text())
			entry.Name = contryList.Text()
		}

		href := s.Find("a")
		if href.Is("a") {
			h, _ := href.Attr("href")
			fmt.Println("link", h)
			entry.Url = h
		}

		font := s.Find("font")
		if font.Is("font") {
			text := font.Text()
			fmt.Println("status", text)
			entry.Status = text
		}

		//fmt.Println(in, " - ", title)
		in++
	})
	fmt.Println("end loop")
	return &entries
}
func removeGarbage(input string) string {
	return strings.Replace(strings.Replace(input, "\t", "", -1), "\n", "", -1)
}
