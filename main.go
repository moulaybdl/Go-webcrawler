package main

// https://blog.boot.dev/path

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 1 {
		fmt.Println("No arguments were provided")
		return
	}


	fmt.Println("CRAWLING IS STARTING ...")
	pages := crawlPage(os.Args[1], os.Args[1],make(map[string]int))

	for url, num := range pages {
		fmt.Println(url, "occured: ", num)
	}




}

