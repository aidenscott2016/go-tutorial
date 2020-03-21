package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, error := http.Get("http://neverssl.com")
	if error != nil {
		fmt.Println("error: ", error)
		os.Exit(1)
	}

	body := make([]byte, 9999)
	n, err := resp.Body.Read(body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v, %d", string(body), n)
}
