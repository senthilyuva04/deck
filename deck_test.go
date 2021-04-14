package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 48 {
		t.Errorf("Error: Expected card length of 48,found %v", len(deck))
	}

	if !strings.EqualFold(deck[0], "Ace Of Spades") {
		t.Errorf("Error:First card is not Ace Of Spades,It is %v", deck[0])
	}
	if !strings.EqualFold(deck[len(deck)-1], "King Of Clubs") {
		t.Errorf("Error:First card is not King Of Clubs,It is %v", deck[len(deck)-1])
	}

}

func TestSaveToDeckAndnewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.SaveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 48 {
		t.Errorf("Error: Expected card length of 48,found %v", len(deck))
	}

	os.Remove("_deckTesting")

}
