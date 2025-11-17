package main

import "math/rand"

func main() {
	cards := newDeckFromFile("list_of_cards")
	cards.print()
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
