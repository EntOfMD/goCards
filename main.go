package main

import "fmt"

func main() {

	cards := deck{newCard(), newCard()}

	for _, i := range cards {
		fmt.Println(i)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
