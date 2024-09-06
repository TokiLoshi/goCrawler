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
	var traverseNode func(*html.Node) 
	traverseNode = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key == "href" {
					parsedURL, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("in getURLsFromHTML '%v': %v\n", anchor.Val, err)
						continue
					}
					fullURL := baseURL.ResolveReference(parsedURL).String()
					urls = append(urls, fullURL)
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNode(child)
		}
	}
	traverseNode(doc)

	if len(urls) == 0 {
		return urls, fmt.Errorf("invalid HTML %v", err)
	}

	return urls, nil
}