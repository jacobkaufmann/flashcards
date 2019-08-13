package flashcards

import (
	"math/rand"
	"time"
)

// Deck represents a deck of cards.
type Deck struct {
	cards []*Card
	index int
}

// NewDeck returns a deck containing cards.
func NewDeck(cards []*Card) *Deck {
	return &Deck{
		cards: cards,
	}
}

// Draw returns the card in the deck at position index.
func (d *Deck) Draw() *Card {
	c := d.cards[d.index]

	d.index++
	if d.index == len(d.cards) {
		d.index = 0
	}

	return c
}

// Shuffle rearranges the cards in a pseudo-random manner and resets the index
// to zero.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	d.index = 0
}
