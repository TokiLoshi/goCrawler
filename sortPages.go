package main

import (
	"sort"
)

type page struct {
	links int
	url string
}

func sortPages(pages map[string]int) []page {

	// sort the pages 
	var sortedPages []page
	for url, links := range pages {
		sortedPages = append(sortedPages, page{url: url, links: links})
	}

	sort.Slice(sortedPages, func(i, j int) bool {
		if sortedPages[i].links == sortedPages[j].links {
			return sortedPages[i].url < sortedPages[j].url
		}
		return sortedPages[i].links > sortedPages[j].links
		})
	
	return sortedPages
}