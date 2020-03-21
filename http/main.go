package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, error := http.Get("http://neverssl.com")
	if error != nil {
		fmt.Println("error: ", error)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("printed %d lines", len(bs))
	return len(bs), nil
}
