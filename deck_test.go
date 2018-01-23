package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T)  {
	d := newDeck()
	expectedCards := 52
	if len(d) != expectedCards {
		t.Errorf("Expected %v cards in a new deck, but there is only %v", expectedCards, len(d))
	}

	expectedFirstCardName := "Ace of Spades"
	if d[0] != expectedFirstCardName{
		t.Errorf("Expected the first card to be %s, but it is %s", expectedFirstCardName, d[0])

	}

	expectedLastCardName := "King of Clubs"
	if d[len(d)-1] != expectedLastCardName{
		t.Errorf("Expected the last card to be %s, but it is %s", expectedLastCardName, d[len(d)-1])

	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T){
	os.Remove("_test")
	deck := newDeck()
	deck.saveToFile("_test")

	restoredDeck := newDeckFromFile("_test")
	if len(restoredDeck) != 52 {
		t.Errorf("Loaded deck must have 52 cards, but there are only %v", len(restoredDeck))
	}
	os.Remove("_test")

}