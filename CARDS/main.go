package main

import "fmt"

func main() {

	//declaring slice
	cards := []string{"Ace of Diamonds", newCard()}
	//adding value to slice
	cards = append(cards, "Six of Spades")
	//iterating through slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

// new function newCard which return string value
func newCard() string {
	return "Aces of Spades"
}
