package main

import "fmt"

// https://blog.boot.dev/path

func main(){
	body := `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`
	r, _ :=getURLsFromHTML(body, "https://Boot.dev")

	fmt.Println(r)
}

