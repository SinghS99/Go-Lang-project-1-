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
	cardSuits := []string{"Heart", "diamond", "club", "joker"}
	cardsValues := []string{"ace", "one", "two", "three"}

	for _, suits := range cardSuits {
		for _, values := range cardsValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//option 1-log the error and return the call to newDeck()
		//option 2-log the error and entirely quit the program if  reading file from hardrive
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPostion := r.Intn(len(d) - 1)

		d[i], d[newPostion] = d[newPostion], d[i]
	}
}
