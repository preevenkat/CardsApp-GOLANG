package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"

	cardStack := newDeck()
	//	cardStack.saveToFile("my_cards")
	//Result := newDeckFromFile("my_cards")
	//	Result.printcards()
	//cardStack.shuffleCards()
		cardStack.printcards()

	//firstSet, lastSet := dealCards(cardStack, 8)
	//firstSet.printcards()
	//lastSet.printcards()
	//fmt.Println(cardStack.toString())

}
