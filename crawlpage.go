package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (map[string]int) {

	baseURLObj, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("error occured parsing the base URL")
		return nil
	}

	currentURLObj, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error occured parsing the current URL: ", rawCurrentURL)
		return pages
	}

	if currentURLObj.Host != baseURLObj.Host {
		return pages
	}

	//! adding the current page to the pages map:
	currentNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return pages
	}

	_, ok := pages[currentNormalized]
	if ok {
		pages[currentNormalized] += 1
		return pages
	} else {
		pages[currentNormalized] = 1
	}

	//! get the html page:
	myHTML, err :=getHTML(rawCurrentURL)
	if err != nil {
		return pages
	}


	//! get the anchors from the html body

	anchors, err := getURLsFromHTML(myHTML, "https://wagslane.dev")
	if err != nil {
		return pages
	}

	for _, link := range anchors {
		pages = crawlPage(rawBaseURL, link, pages)
	}


	return pages

}