package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println("error: ", error)
		os.Exit(1)
	}

	fmt.Println(resp.Header)
}
