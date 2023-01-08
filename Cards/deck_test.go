package main

import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()


	if(len(d) != 6){
		t.Errorf("Expected deck length of 6, but got %d", len(d))
	}

	if (d[0] != "Spards of Ace1") {
		t.Errorf("Expected first card of deck is Spards of Ace, but got %v", d[0])
	}

}
