package main

import "fmt"

func main() {
	//variation in Declaration

	card := newCard() //declaring a new variable as string

	fmt.Println(card)

}

// new function newCard which return string value
func newCard() string {
	return "Five of Diamonds"
}
