package main

func main() {
	cards := newDeck()
	cards.saveToFile("card")
	loadFromFile("card").print()

}
