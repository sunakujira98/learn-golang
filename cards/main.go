package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// cards := newDeckFromFilename("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards.saveToFile("my_cards")
}

func newCard() string {
	return "Five of diamonds"
}
