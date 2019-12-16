package main
// package main is the executable file and it calls the main function
import (
	"math/rand"
	"fmt"
	"time"
)
func main(){
	cardMaker()
	randomNum()
	fmt.Println(pickACard(randomNum(), cardMaker()))
}
func pickACard(id int, deck []string)string {
	return deck[id]
}

func cardMaker() []string{
	var box []string
	suites := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	names := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	deck := len(suites)
	cards := len(names)
	for index := 0; index < deck; index++ {
		for j := 0; j < cards; j++ {
		  card := names[j] + " of " + suites[index]
		  box = append(box, card)
		}
	}
	return box
}

func randomNum() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(52)
}