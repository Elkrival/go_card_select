package main

import ("os"
"testing")
func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Got Ace of spades %v", d[0])
	}
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Did not get king of diamonds instead got %v", d[0])
	}
}

// Os has the remove function which will remove files from the hard drive of the computer
// * 
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting");
	loadedDeck := newDeckFromFile("_decktesting") 
	if len(loadedDeck) != 52 {
		t.Errorf("This deck does not have a length of 52. %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
