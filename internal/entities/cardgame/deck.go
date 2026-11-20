package cardgame

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

// Deck rappresenta un mazzo di carte.
type Deck struct {
	cards      []Card
	jokerCount int
}

// NewDeck crea e mescola un mazzo standard di 52 carte.
func NewDeck() *Deck {
	deck := &Deck{cards: make([]Card, 0, 52)}
	for value := 1; value <= 13; value++ {
		for _, suit := range []string{"heart", "diamonds", "spades", "clubs"} {
			deck.cards = append(deck.cards, Card{Value: value, Suit: suit})
		}
	}
	return deck
}

// String restituisce una rappresentazione testuale del mazzo.
func (d *Deck) String() string {
	return fmt.Sprintf("cards: %v", d.cards)
}

// Cards restituisce una copia delle carte nel mazzo.
func (d *Deck) Cards() []Card {
	return append([]Card(nil), d.cards...)//crea una lista di card nulla e gli passa le card dentro cards una a una
}

// CardCount restituisce il numero di carte nel mazzo.
func (d *Deck) CardCount() int {
	return len(d.cards)
}

// Sort ordina le carte prima per seme, poi per valore. TODO bisogna fixarlo non funziona.
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

// AddCard aggiunge una carta al mazzo.
func (d *Deck) AddCard(card Card) {
	d.cards = append(d.cards, card)
}

// Remove rimuove la carta indicata e segnala se è stata trovata.
func (d *Deck) Remove(value int, suit string) bool {
	for i, card := range d.cards {
		if card.Value == value && card.Suit == suit {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
			return true
		}
	}
	return false
}

// DrawSpecificCard estrae la carta richiesta oppure restituisce un errore.
func (d *Deck) DrawSpecificCard(value int, suit string) (Card, error) {
	for i, card := range d.cards {
		if card.Value == value && card.Suit == suit {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
			return card, nil
		}
	}
	return Card{}, errors.New("card not in deck")
}

// PopFirst estrae la prima carta, o restituisce un errore se il mazzo è vuoto.
func (d *Deck) PopFirst() (Card, error) {
	if len(d.cards) == 0 {
		return Card{}, errors.New("deck is empty")
	}
	card := d.cards[0]
	d.cards[0] = Card{}
	d.cards = d.cards[1:]
	return card, nil
}

// Shuffle mescola le carte.
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// CreateJoker crea al massimo due jolly.
func (d *Deck) CreateJoker() (Card, error) {
	if d.jokerCount >= 2 {
		return Card{}, errors.New("all jokers already created")
	}else if d.jokerCount == 1{
		return Card{Value: 0, Suit: "spades"},nil
	}
	d.jokerCount++
	return Card{Value: 0, Suit: "heart"}, nil
}

// AddJoker crea un jolly e lo aggiunge al mazzo.
func (d *Deck) AddJoker() error {
	joker, err := d.CreateJoker()
	if err != nil {
		return err
	}
	d.AddCard(joker)
	return nil
}
