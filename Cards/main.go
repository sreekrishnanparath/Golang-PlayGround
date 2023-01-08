package main

import (
	"fmt"
)

func main() {
	//print()

	//cards := newDeck()
	//hand, remainDeck := cards.chooseDecks(1)
	//hand.print()
	//remainDeck.print()

	hello := "Hello"
	fmt.Println([]byte(hello))
	//fmt.Println(cards.toString())
	//cards.saveToFile("D:/gotestfile.txt")
	cards := readFromFile("D:/gotestfile.txt")
	cards.print()
}
