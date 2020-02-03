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

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d deck is receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// no receiver down here
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// byte slice and error object
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// time and rand pkgs
	source := rand.NewSource(time.Now().UnixNano())
	// use source to generate new rand num generator, r, to randomize deck
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swap
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
