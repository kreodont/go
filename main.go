package main

func main()  {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("shuffled")

}


