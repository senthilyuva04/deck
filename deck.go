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
	newDeck := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, value+" Of "+suit)
		}
	}
	return newDeck
}

func (d deck) Print() {
	for i, deck := range d {
		fmt.Println(i, deck)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) SaveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.ToString()), 0666)
}
func newDeckFromFile(fileName string) deck {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(b), ","))
}

func (d deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		rn := r.Intn(len(d) - 1)
		d[i], d[rn] = d[rn], d[i]
	}

}
