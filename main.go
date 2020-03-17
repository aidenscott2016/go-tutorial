package main

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 3)
	cards.saveToFile("cards")

}
