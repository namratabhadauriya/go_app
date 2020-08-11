package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"heart", "diamond", "spade", "club"}
	cardValues := []string{"one", "two", "three", "four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards

}

func deal(d deck, handnumber int) (deck, deck) {
	return d[:handnumber], d[handnumber:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffleCards() {
	for i, _ := range d {
		position := rand.Intn(len(d) - 1)
		d[i], d[position] = d[position], d[i]
	}
}
