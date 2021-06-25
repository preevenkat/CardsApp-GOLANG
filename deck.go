package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//new deck type with slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suites := range cardSuites {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suites)

		}
	}
	return cards
}

func (d deck) printcards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func dealCards(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}
func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(cardsFile string) error {
	return ioutil.WriteFile(cardsFile, []byte(d.toString()), 0666)

}

func newDeckFromFile(fileName string) deck {

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//ss is slice of string
	ss := strings.Split(string(bs), ",")
	return deck(ss)

}

func (d deck) shuffleCards() {
	sourceValue := rand.NewSource(time.Now().Unix())
	rGen := rand.New(sourceValue)
	for index := range d {
		newPostion := rGen.Intn(len(d) - 1)
		d[index], d[newPostion] = d[newPostion], d[index]

	}
}
