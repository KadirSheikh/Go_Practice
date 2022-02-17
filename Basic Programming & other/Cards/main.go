package main

import (
	"fmt"
	"os"
)

func main() {

	//hand, remainingDeck := deal(cards, 5)

	// hand.printDeck()
	// fmt.Println("---------------------------------------------------")
	// remainingDeck.printDeck()
	// fmt.Println("---------------------------------------------------")

	// data := cards.toStringToByteSlice()

	// fmt.Println(data)

	// cards.saveToFile("cards.txt")
	// hand.saveToFile("hand.txt")
	// remainingDeck.saveToFile("remainingDeck.txt")

	// byteData, error := newDeckFromFile("cards.txt")

	// fmt.Println(byteData, error)

	// cards := newDeckFromFile("cards.txt")
	// cards.shuffle()
	// cards.printDeck()

	fmt.Println("**********MENU***********")
	fmt.Println("1.Add New Deck")
	fmt.Println("2.Print Deck")
	fmt.Println("3.Make Deal")
	fmt.Println("4.Print Hand")
	fmt.Println("5.Print Remaining Deck")
	fmt.Println("6.Shuffle Deck")
	fmt.Println("7.Save Deck to File")
	fmt.Println("8.Read Deck from File")
	fmt.Println("9.Exit")

	for {
		cards()
	}

}

func cards() {
	var choice int

	fmt.Println("Enter Your Choice : ")
	fmt.Scanf("%d", &choice)

	cards := deck{}
	hand := deck{}
	remainingDeck := deck{}

	switch choice {
	case 1:

		cards = newDeck()

		if len(cards) == 40 {
			fmt.Println("Deck Created")
		}

	case 2:
		fmt.Println("2")
		cards.printDeck()

	case 3:
		var i int
		fmt.Scanf("%d", &i)
		fmt.Println("Enter Hand Size : ")
		d, _ := fmt.Scanf("%d", &i)
		hand, remainingDeck = deal(cards, d)
	case 4:
		hand.printDeck()
	case 5:
		remainingDeck.printDeck()
	case 6:
		cards.shuffle()
	case 7:
		cards.saveToFile("cards.txt")
	case 8:
		cards = newDeckFromFile("cards.txt")
	case 9:
		os.Exit(1)
	default:

		fmt.Println("Invalid Choice")
	}

}
