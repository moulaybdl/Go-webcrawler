package main

import (
	"errors"
	"fmt"
	"net/url"
)

func normalizeURL(rawURL string) (string, error) {
	urlObj, err := url.Parse(rawURL)
	if err != nil {
			return "", err
	}
	
	if urlObj.Host == "" {
		return "", errors.New("no Hostname provided")
	}

	pathPart := ""

	if urlObj.Path != "" {
		if string(urlObj.Path[len(urlObj.Path) - 1]) == "/" {
			pathPart = urlObj.Path[:len(urlObj.Path) - 1]
		} else {
			pathPart = urlObj.Path
		}
	}

	normalizedURL := fmt.Sprintf("%s%s", urlObj.Host, pathPart)

	return normalizedURL, nil
}