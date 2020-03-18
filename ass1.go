package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
