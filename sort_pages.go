package main

import (
	"sort"
)

type Page struct {
	links int
	url string
}

func sortPages(pages map[string]int) []Page {

	// sort the pages 
	var sortedPages []Page
	for url, links := range pages {
		sortedPages = append(sortedPages, Page{url: url, links: links})
	}

	sort.Slice(sortedPages, func(i, j int) bool {
		if sortedPages[i].links == sortedPages[j].links {
			return sortedPages[i].url < sortedPages[j].url
		}
		return sortedPages[i].links > sortedPages[j].links
		})
	
	return sortedPages
}