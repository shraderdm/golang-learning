package main

func main() {
	cards := newDeckFromFile("my_cards")
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards.shuffle()
	cards.print()
}
