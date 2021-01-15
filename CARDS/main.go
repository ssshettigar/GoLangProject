package main

import "fmt"

var deckSize int //variable can be declared outside and you can't assign value here

func main() {
	//variation in Declaration
	//var card string = "Ace of Spades"
	card := "Ace of Spades"  //declaring a new variable as string
	card = "Ace of Diamonds" //assigning a new value to card variable
	fmt.Println(card)

	deckSize = 50
	fmt.Println(deckSize)
}
