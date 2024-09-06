package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	// Initialize the concurrency control
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	// Check rawCurrentURL is on same domain as rawBaseURL
	parsedBaseURL := cfg.baseURL
	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil { 
		fmt.Printf("Error: crawlPage could not parse current url %v \n", err)
		return 
	}

	// Check the domains 
	baseDomainName := parsedBaseURL.Hostname()
	currentDomainName := parsedCurrent.Hostname()
	
	// Skip other websites
	if baseDomainName != currentDomainName {
		return
	}

	// Get normalized version of rawCurrentURL 
	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("issue normalizing rawCurrentURL: %v", err)
		return 
	}
	
	
	// Handle Base case
	isFirst := cfg.addPageVisit(normalizedCurrent)
	if !isFirst {
		return
	}

	// Get HTML from current URL and add print statement to watch crawler 
	fmt.Printf("On Crawler: crawling: %s\n", rawCurrentURL)

	
	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("issue getting html: %v", err)
		return 
	}

	// Get URLS
	allURLs, err := getURLsFromHTML(currentHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("Issue getting all the urls: %v", err)
	}
	fmt.Println("I think we got all the URLS!")
	
	for i := 0; i < len(allURLs); i ++ {
		fmt.Println(allURLs[i])
	}
	
	// Recursively crawl each url on the page 
	for _, webURL := range allURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(webURL)
	}

}