package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(HTMLbody, rawURL string)([]string ,error){
	var anchors []string

	r := strings.NewReader(HTMLbody)
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}


	var f func(*html.Node)

	f = func(el *html.Node){
		if el.Type == html.ElementNode && el.Data == "a" {
			attribute := el.Attr
			for _, attr := range attribute {
				if attr.Key == "href" {
					urlObj, err := url.Parse(attr.Val)
					if err != nil {
						break
					}

					if urlObj.Host == "" {
						anchors = append(anchors, rawURL+attr.Val)
					} else {
						anchors = append(anchors, attr.Val)	
					}
				}
			}

		}

		for c := el.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(root)

	return anchors, nil
}