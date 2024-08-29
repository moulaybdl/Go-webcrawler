package main

import (
	"fmt"
)

// https://blog.boot.dev/path

func main(){
	nromalizedURL, err := normalizeURL("#")
	if err != nil {
		fmt.Println("error from normalized function")
		return
	}
	fmt.Println("the normalized url:", nromalizedURL)
}

