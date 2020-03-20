package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type deck
// which is a slice of strings
type  deck []string



func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value+" of "+suit
			cards = append(cards,card)
		}
	}

	return  cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}


func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}


func (d deck) toString() string {
 return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck  {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}


func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newI := r.Intn(len(d) - 1)

		d[i], d[newI] = d[newI], d[i]
	}
}