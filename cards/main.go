package main

func main() {
	cards := newDeck()
	// cards := newDeckFromFile("./hand.txt")
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)
	// fmt.Println("hand")
	// fmt.Println("toString: " + hand.toString())
	// hand.saveToFile("./hand.txt")
	// hand.print()
	// fmt.Println("RemainingCards")
	// remainingCards.print()

	// greeting := "こんにちは"
	// fmt.Println([]byte(greeting))
}
