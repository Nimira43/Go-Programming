package main

func main() {
	cards := newDeckFromFile("list_of_cards")
	cards.print()	
}

func (d deck) shuffle() {

}
