package main

import (
	"fmt"
	"time"
)

func main() {
	input := []int{4, 2, 6, 7, 1, 3, 9, 23}
	output := []int{}

	c := make(chan int)
	for _, d := range input {

		go func(t int, c chan int) {
			time.Sleep(time.Millisecond * time.Duration(t))
			c <- t
		}(d, c)
	}

	for i := 0; i < len(input); i++ {
		output = append(output, <-c)
	}

	fmt.Println(output)
}
