package main

func main() {
	cards := newDeck()
	hand, remainingCards, other := deal(cards, 5)
	hand.print()
	remainingCards.print()
	other.print()
}
