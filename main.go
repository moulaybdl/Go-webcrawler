package main

import (
	"fmt"
	"net/url"
)

// https://blog.boot.dev/path

func main(){
	urlObj, err := url.Parse("https://")
	if err != nil {
		fmt.Println("Smth went wrong with the parsing: ", err)
		return
	}

	fmt.Printf("the return is: %v",urlObj.Scheme )

}

