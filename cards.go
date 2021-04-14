package main

func main() {
	// deck := newDeck()
	// fmt.Println(deck.toString())
	// deck.saveToFile("test")
	deck := newDeckFromFile("test")
	deck.Print()
	deck.Shuffle()
	deck.Print()
}
