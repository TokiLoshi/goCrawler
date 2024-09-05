package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)


func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	
	urls := []string{}

	if len(rawBaseURL) == 0 {
		return urls, fmt.Errorf("empty base URL")
	}
	
	// Parse the base url 
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil { 
		return urls, fmt.Errorf("error parsing URL: %v", err)
	}
	
	// Parse the HTML 
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return urls, fmt.Errorf("invalid HTML: %v", err)
	}

	
	// create a tree of html nodes
	
	// recursively traverse the treee node and append the anchor tags
	var f func(*html.Node) 
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					parsedURL, err := url.Parse(a.Val)
					if err != nil {
						continue
					}
					fullURL := baseURL.ResolveReference(parsedURL).String()
					urls = append(urls, fullURL)
				}
			}
		}
		for c:= n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if len(urls) == 0 {
		return urls, fmt.Errorf("invalid HTML %v", err)
	}

	return urls, nil
}