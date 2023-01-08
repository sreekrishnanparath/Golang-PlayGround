package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spards","Clubs"}
	cardValue := []string{"Ace","One","Two"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, suit+ " of " +value)
		}
	}
	return cards
}

// func chooseDecks(d deck, handSize int) (deck, deck) {
// 	return d[:handSize], d[handSize:]
// }

func (d deck) chooseDecks(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([] string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	bs, error := ioutil.ReadFile(filename)
	if(error != nil){
		fmt.Println(error)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}