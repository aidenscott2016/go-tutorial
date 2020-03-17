package main

import "fmt"

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 3)
	fmt.Println(cards.toString())

}
