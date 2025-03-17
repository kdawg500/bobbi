package crawl

import (
	"net/http"

	"golang.org/x/net/html"
)

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

func extractDocHREFs(doc *html.Node, linkMap map[string]int) {
	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, attribute := range doc.Attr {
			if attribute.Key == "href" {
				linkMap[attribute.Val]++
				break
			}
		}
	}
	for d := doc.FirstChild; d != nil; d = d.NextSibling {
		extractDocHREFs(d, linkMap)
	}
}
