package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// use http.get to fetch webpage of rawURL
	
	// METHOD 1:
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	statusCode := res.StatusCode
	fmt.Println("StatusCode: ", statusCode)

	// If HTTP status code is 400+ return err
	if statusCode >= 400 {
		fmt.Println("Errorin the 400s: ", statusCode)
		return "", fmt.Errorf("error 400+: %d", statusCode)
	}
	if statusCode >= 300 && statusCode < 400 {
		fmt.Println("Error in the 300s", statusCode)
		return "", fmt.Errorf("error in 300s: %d", statusCode)
	}
	// If content-type-header is not text/html return error
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("content type is not text/html: %s", contentType)
	}
	
	// Read html body 
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("something went wrong returning page contents %v", err)
	}
	bodyString := string(bodyBytes)
	fmt.Printf("Reader returns: %s", bodyString)

	fmt.Println("Body: ", bodyString)
	return bodyString, nil

	
	
	// return any other possible errors
	// return webpage's html if all goes well
	// use io.ReadAll to read the response 
	// call getHTML from main 
	// print the result (some html from the internet)
}