package link

import (
	"net/url"
	"strings"
)

func Validate(link string) string {
	if !strings.HasPrefix(link, "https://") {
		link = "https://" + link
	}

	parsedURL, err := url.Parse(link)
	if err != nil {
		panic(err)
	}

	hostSegments := strings.Split(parsedURL.Host, ".")
	if len(hostSegments) == 1 {
		panic("bobbi: invalid url provided, missing domain")
	}

	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	}

	return parsedURL.String()
}
