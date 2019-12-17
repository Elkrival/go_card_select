package main
// package main is the executable file and it calls the main function
import (
	// "time"
)
// func main(){
// 	cardMaker()
// 	randomNum()
// 	fmt.Println(pickACard(randomNum(), cardMaker()))
// }
// func pickACard(id int, deck []string)string {
// 	return deck[id]
// }

// func cardMaker() []string{
// 	var box []string
// 	suites := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
// 	names := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
// 	deck := len(suites)
// 	cards := len(names)
// 	for index := 0; index < deck; index++ {
// 		for j := 0; j < cards; j++ {
// 		  card := names[j] + " of " + suites[index]
// 		  box = append(box, card) //Append does not add the element to the current slice it returns a new slice with all the elements and assigns it to the variable
// 		}
// 	}
// 	return box
// }

// func randomNum() int{
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(52)
// }

func main (){
	cards := newDeck()
	//cards is of type deck and it will have access to the print method
	/*
		Iterating over a closed set
	 for: loop syntax. First variable will always lead to index, second variable will be value
	 index: index of the element in slice
	 card: current item that i am iterating
	 range cards: take the slice "cards" and loop over it
	for index, card := range cards
	*/
	// The deal function will return two values and to assign those values correctly
	// We will assign them to two variables
	// cards.saveToFile("decks_cards")
	// newDeck := newDeckFromFile("decks_cards")
	// newDeck.print()
	cards.shuffle()
	cards.print()
}
