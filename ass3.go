package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args[1]) > 0 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, f)
	} else {
		fmt.Println("missing file param")
		os.Exit(1)
	}

}
