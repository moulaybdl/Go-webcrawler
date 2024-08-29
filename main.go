package main

import "fmt"

// https://blog.boot.dev/path

func main(){
	r, err := normalizeURL("https://bLog.BoOt.dEv/paTh")
	if err != nil {
		fmt.Println("Error ya 7mar", err)
		return
	}

	fmt.Println(r)
}

