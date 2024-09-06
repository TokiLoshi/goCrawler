package main

import (
	"fmt"
	"net/url"
	"sync"
)


type config struct {
	// Keep track of pages we've crawled
	pages map[string]int
	// Keep track of original base URL
	baseURL *url.URL
	// Ensure pages is threadsafe 
	mu *sync.Mutex 
	// Ensuring we don't spawn too many goroutines at once
	// Buffered channel of empty structs 
	// When new goroutine starts send empty struct to channel
	// When it's done receive empty struct from channel
	// this will block new go routines and wait until buffer has space
	// eg buffer size of 5 means 5 requests at once 
	concurrencyControl chan struct{}
	// Ensure main waits until all in-flight goroutines (HTTP requests) are done
	// Then exit 
	wg *sync.WaitGroup
}


// Helper function to track pages 
func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	// add a mutex lock 
	cfg.mu.Lock()
	// defer the unlock 
	defer cfg.mu.Unlock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}

func configure(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}
	return &config {
		pages: make(map[string]int),
		baseURL: baseURL, 
		mu: &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency), 
		wg: &sync.WaitGroup{},
	}, nil
}