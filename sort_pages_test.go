package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name string
		input map[string]int 
		expected []Page
	}{
		{
		name: "order count decending",
		input: map[string]int{
			"url1" : 5, 
			"url2" : 1,
			"url3" : 3,
			"url4" : 10,
			"url5" : 7,
		},
		expected: []Page{
			{url: "url4", links: 10}, 
			{url: "url5", links: 7},
			{url: "url1", links: 5},
			{url: "url3", links: 3},
			{url: "url2", links: 1},
		},
	},
}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.input)
			if !reflect.DeepEqual(actual, tc.expected){
				t.Errorf("Text %v - %s FAIL: expected %v, got: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}



