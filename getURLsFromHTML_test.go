package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

// writing relative URLs and passing them to an array
// the array is un-normalized strings  of URLs, or errors
// refelect.DeepEqual function will be helpful
// need to test that urls are converted to absolute urls
// find all the <a>tags in the html body

func TestURLsFromHTML(t * testing.T) {

	tests := []struct {
		name string 
		rawBaseURL string
		htmlBody string 
		expected []string
		errorContains string
	}{
    { 
		name:     "absolute and relative URLs",
    rawBaseURL: "https://blog.boot.dev",
    htmlBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/two">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
    expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/two"},

		},
		{
			name: "all the a tags are collected",
			rawBaseURL: "https://blog.boot.dev",
			htmlBody: `<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/example">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/two">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`, 
expected: []string{"https://blog.boot.dev/path/one", "https://other.com/example", "https://other.com/path/two"},
		},
		{
			name: "urls are converted to absolute urls",
			rawBaseURL: "http://blog.golang.com",
			htmlBody: `<html>
	<body>
		<a href="/absolute-gopher">
			<span>GoPher.dev</span>
		</a>
	</body>
</html>
`, 
expected: []string{"http://blog.golang.com/absolute-gopher"},
		},
		// {
		// 	name: "invalid HTML",
		// 	rawBaseURL: "https://blog.boot.dev",
		// 	htmlBody: `<html><body><a href=">malformed link</a></body></html>`,
		// 	expected: nil, 
		// 	errorContains: "error parsing HTML",
		// },
		// {
		// 	name: "empty base URL", 
		// 	rawBaseURL: "", 
		// 	htmlBody: `<html><body><a href="/relative>relative link</a></body></html>`,
		// 	expected: nil, 
		// 	errorContains: "error parsing url",
		// },
	
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.rawBaseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}
			actual, err := getURLsFromHTML(tc.rawBaseURL, baseURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s FAIL: expercted absolute path '%v' and got '%v", i, tc.name, tc.expected, actual)
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v' got none", i, tc.name, tc.errorContains)
			}
			// reflect DeepEqual
			//https://cs.opensource.google/go/go/+/refs/tags/go1.23.0:src/reflect/deepequal.go
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: %v, actual %v", i, tc.name, tc.expected, actual)
				return
			}
			
		})
	}

}