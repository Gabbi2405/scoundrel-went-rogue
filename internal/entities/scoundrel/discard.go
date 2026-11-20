package scoundrel

import (
	"fmt"
	"scoundrelWentRogue/internal/entities/cardgame"
	"sort"
	"strings"
)

// DiscardPile holds cards discarded during the game.
type DiscardPile struct {
	cards []cardgame.Card
}

// NewDiscardPile creates an empty discard pile.
func NewDiscardPile() *DiscardPile {
	return &DiscardPile{cards: make([]cardgame.Card, 0)}
}

// AddCard adds a card to the discard pile.
func (p *DiscardPile) AddCard(card cardgame.Card) {
	p.cards = append(p.cards, card)
}

// Cards returns a copy of the discarded cards.
func (p *DiscardPile) Cards() []cardgame.Card {
	return append([]cardgame.Card(nil), p.cards...)
}

// CardCount returns the number of discarded cards.
func (p *DiscardPile) CardCount() int {
	return len(p.cards)
}

// Sort orders by suit and value without modifying the pile order.
func (p *DiscardPile) Sort() {
	order := map[string]int{"heart": 0, "diamonds": 1, "spades": 2, "clubs": 3}
	sort.Slice(p.cards, func(i, j int) bool {
		si, okI := order[p.cards[i].Suit]
		sj, okJ := order[p.cards[j].Suit]
		if okI && okJ {
			// Both suits in order map: compare by order
			if si != sj {
				return si < sj
			}
			return p.cards[i].Value < p.cards[j].Value
		}
		if okI || okJ {
			// Only one in order map: known suits come first
			return okI
		}
		// Neither in map: compare alphabetically
		if p.cards[i].Suit != p.cards[j].Suit {
			return p.cards[i].Suit < p.cards[j].Suit
		}
		return p.cards[i].Value < p.cards[j].Value
	})
}

// String returns a readable representation of the room.
func (r *Room) String() string {
	var result strings.Builder
	for i, card := range r.cards {
		if i > 0 {
			result.WriteByte('\n')
		}
		fmt.Fprint(&result, card)
	}
	return result.String()
}

