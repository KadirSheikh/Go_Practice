package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuit := deck{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValue := deck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			// fmt.Println(i, j, suit, value)
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) toStringToByteSlice() []byte {

	return []byte(strings.Join([]string(d), ","))

}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, d.toStringToByteSlice(), 0666)

}

func newDeckFromFile(filename string) deck {

	bs, error := ioutil.ReadFile(filename)

	if error != nil {
		fmt.Println("Error reading:", error)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {

		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}

}
