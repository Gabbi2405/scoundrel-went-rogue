package cardgame

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

// Deck represents a deck of cards.
type Deck struct {
	cards      []Card
	jokerCount int
}

// NewDeck creates and shuffles a standard 52-card deck.
func NewDeck() *Deck {
	deck := &Deck{cards: make([]Card, 0, 52)}
	for value := 1; value <= 13; value++ {
		for _, suit := range []string{"heart", "diamonds", "spades", "clubs"} {
			deck.cards = append(deck.cards, Card{Value: value, Suit: suit})
		}
	}
	return deck
}

// String returns a textual representation of the deck.
func (d *Deck) String() string {
	return fmt.Sprintf("cards: %v", d.cards)
}

// Cards returns a copy of the cards in the deck.
func (d *Deck) Cards() []Card {
	return append([]Card(nil), d.cards...)// creates a nil slice and copies cards into it one by one
}

// CardCount returns the number of cards in the deck.
func (d *Deck) CardCount() int {
	return len(d.cards)
}

// Sort orders cards by suit then value.
func (d *Deck) Sort() {
	order := map[string]int{"heart": 0, "diamonds": 1, "spades": 2, "clubs": 3}
	sort.Slice(d.cards, func(i, j int) bool {
		si, okI := order[d.cards[i].Suit]
		sj, okJ := order[d.cards[j].Suit]
		if okI != okJ {
			return okI
		}
		if !okI {
			return d.cards[i].Suit < d.cards[j].Suit
		}
		if si == sj {
			return d.cards[i].Value < d.cards[j].Value
		}
		return si < sj
	})
}

// AddCard adds a card to the deck.
func (d *Deck) AddCard(card Card) {
	d.cards = append(d.cards, card)
}

// Remove removes the specified card and reports whether it was found.
func (d *Deck) Remove(value int, suit string) bool {
	for i, card := range d.cards {
		if card.Value == value && card.Suit == suit {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
			return true
		}
	}
	return false
}

// DrawSpecificCard draws the requested card or returns an error.
func (d *Deck) DrawSpecificCard(value int, suit string) (Card, error) {
	for i, card := range d.cards {
		if card.Value == value && card.Suit == suit {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
			return card, nil
		}
	}
	return Card{}, errors.New("card not in deck")
}

// PopFirst draws the first card, or returns an error if the deck is empty.
func (d *Deck) PopFirst() (Card, error) {
	if len(d.cards) == 0 {
		return Card{}, errors.New("deck is empty")
	}
	card := d.cards[0]
	d.cards[0] = Card{}
	d.cards = d.cards[1:]
	return card, nil
}

// Shuffle shuffles the cards.
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// CreateJoker creates at most two jokers.
func (d *Deck) CreateJoker() (Card, error) {
	if d.jokerCount >= 2 {
		return Card{}, errors.New("all jokers already created")
	}else if d.jokerCount == 1{
		return Card{Value: 0, Suit: "spades"},nil
	}
	d.jokerCount++
	return Card{Value: 0, Suit: "heart"}, nil
}

// AddJoker creates a joker and adds it to the deck.
func (d *Deck) AddJoker() error {
	joker, err := d.CreateJoker()
	if err != nil {
		return err
	}
	d.AddCard(joker)
	return nil
}
