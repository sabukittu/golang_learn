package main

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 3)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	cards.shuffle()
	cards.print()

}
