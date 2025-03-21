package crawl

import (
	"net/http"

	"golang.org/x/net/html"
)

type Page struct {
	URL      string
	doc      *html.Node
	children map[string]Page
}

func (p *Page) GetHTMLDoc() {

}

func (p *Page) insertChildren() error {
	doc, err := getHTMLDoc(p.URL)
	if err != nil {
		return err
	}

	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, attribute := range doc.Attr {
			if attribute.Key == "href" {
				p.children[attribute.Val] = Page{}
				break
			}
		}
	}
	for d := doc.FirstChild; d != nil; d = d.NextSibling {
		extractDocHREFs(d)
	}
	return nil
}

func (p *Page) Crawl() {

}

func getHTMLDoc(link string) (*html.Node, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func extractDocHREFs(doc *html.Node) map[string]int {
	linksMap := map[string]int{}
	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, attribute := range doc.Attr {
			if attribute.Key == "href" {
				linksMap[attribute.Val]++
				break
			}
		}
	}
	for d := doc.FirstChild; d != nil; d = d.NextSibling {
		extractDocHREFs(d)
	}
	return linksMap
}
