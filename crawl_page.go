package main

import (
	"fmt"
	"net/url"
)


func (cfg *config) crawlPage(rawCurrentURL string) {
	// Check to see if we've maxed out pages
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

	if !cfg.addPageVisit(normalizedCurrent) {
		return
	}

	// Get HTML from current URL and add print statement to watch crawler 
	fmt.Printf("crawling: %s\n", rawCurrentURL)

	
	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("issue getting html: %v\n", err)
		return 
	}

	// Get URLS
	allURLs, err := getURLsFromHTML(currentHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("Issue getting all the urls: %v\n", err)
	}
	
	// for i := 0; i < len(allURLs); i ++ {
	// 	fmt.Println(allURLs[i])
	// }
	
	// Recursively crawl each url on the page 
	for _, webURL := range allURLs {
		cfg.mu.Lock()
		if len(cfg.pages) >= cfg.maxPages {
			cfg.mu.Unlock()
			return
		}
		cfg.mu.Unlock()
		cfg.wg.Add(1)
		go cfg.crawlPage(webURL)
	}

}