package link

import (
	"net/url"
)

func Validate(link string) string {
	parsedURL, err := url.Parse(link)
	if err != nil {
		panic(err)
	}

	return parsedURL.String()
}
