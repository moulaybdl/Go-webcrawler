package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string)(string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		fmt.Println("Error occured while fetching the html")
		return "", err
	}
	defer resp.Body.Close() // Ensure the response body is closed


	if resp.StatusCode > 399 {
		return "",errors.New("something went wrong with the request")
	}

	if resp.Header.Get("Content-Type") != "text/html" {
		return "",errors.New("wrong type retured") 
	}

	rawHTML, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading the html")
		return "", nil
	}

	return string(rawHTML), nil
}
