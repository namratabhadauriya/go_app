package main

func main() {
	//cards := newDeck()
	//hand, remaining := deal(cards, 5)
	// hand.print()
	// remaining.print()
	//fmt.Println(cards.toString())
	//cards.saveFile("my_cards")
	cards := readFile("my_cards")
	cards.print()
	cards.shuffleCards()
	cards.print()
}
