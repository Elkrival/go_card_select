package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

// Creates a new type of deck wchi is a slice of strengnts
/*
 Receiver: used when a function is written without a name
	(d, deck) = receiver this means that any variable that is assigned
	the type of deck will have this method available to it
	d = this is the reference of the instance of the deck variable
	for example cards.print() on the main.go file
	cards is our deck type variable and it will be accessed as the variable d of our receiver on the print method
	deck= reference to the type we want attach the method to

	in syntax the convention for a receiver is using the letter abbrev


*/

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	// use of _ tells go that the variable will not be used
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+ suit)
		}
	}
	fmt.Println(cards)
	return cards
}

func (d deck) print() deck {
	for i, card := range d {
		fmt.Println(i, card)
	}
	return d
}
//deal function will take an existing deck of cards and will make a new deck of cards without the cards that were dealt
// This will be splitting of the slice for two other slices
//to select a range of a slice we use the following syntax deck[startIndexIncluding: endIndexNotIncluding]
// Go can return multiple values from a function noted by (deck, deck) 

func deal(d deck, handSize int) (deck, deck){

 return d[:handSize], d[handSize:]
 //this returns :handSize, everything from the start of the slice to the index handSize
 //this returns everything from the handSize to the end
}
/*
	[]byte
	it is a slice of ascii character code of a value
	How to convert a type slice to a []byte slice
	There is a process called type conversion a type of value and converting
	
*/

func (d deck)toString() string{
	//Converting a type to a string that we will use to write to a file
	return strings.Join([]string(d), ",")
}

func (d deck)saveToFile(filename string) error{
	/*
	ioutil: This package handles io input and has methods such as writeToFile or ReadFromfile
	In this instant we are using this to write the file and read the file with the deck of cards
	that we have created
	*/
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
} 
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Handle Error here
		// Optoin 1: log the erro and return a call to newDeck()
		// Option 2: Log the error and quit the program to quit the program we will use the os.Exit method
		fmt.Println(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	// We can do this because we are only using a slice of strings
	return deck(s)

}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// this is a single line reassignment of the decks position and new position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}