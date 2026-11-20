package scoundrel

import (
	"fmt"
	"strings"

	"scoundrelWentRogue/internal/entities/cardgame"
)

// Hand holds the equipped weapon and monsters defeated with that weapon.
type Hand struct {
	weapon *cardgame.Card
	cards  []cardgame.Card
}

// NewHand creates an empty hand.
func NewHand() *Hand {
	return &Hand{cards: make([]cardgame.Card, 0)}
}

// Equip equips a weapon. If a weapon is already equipped, it discards it along with defeated monsters.
func (h *Hand) Equip(card cardgame.Card, discard *DiscardPile) {
	if h.weapon != nil {
		for _, defeated := range h.cards {
			discard.AddCard(defeated)
		}
		discard.AddCard(*h.weapon)
		h.cards = h.cards[:0]
	}
	weapon := card
	h.weapon = &weapon
}

// AddCard records a monster defeated with the currently equipped weapon.
func (h *Hand) AddCard(card cardgame.Card) {
	h.cards = append(h.cards, card)
}

// Weapon returns the equipped weapon and whether one is present.
func (h *Hand) Weapon() (cardgame.Card, bool) {
	if h.weapon == nil {
		return cardgame.Card{}, false
	}
	return *h.weapon, true
}

// DefeatedCards returns a copy of the defeated monsters.
func (h *Hand) DefeatedCards() []cardgame.Card {
	return append([]cardgame.Card(nil), h.cards...)
}

// CardCount returns the number of monsters defeated with the current weapon.
func (h *Hand) CardCount() int {
	return len(h.cards)
}

// String returns a readable representation of the hand.
func (h *Hand) String() string {
	var result strings.Builder
	if h.weapon == nil {
		result.WriteString("weapon equipped: none\n")
	} else {
		fmt.Fprintf(&result, "weapon equipped: %s\n", h.weapon)
	}
	result.WriteString("monster(s) defeated:")
	for _, card := range h.cards {
		fmt.Fprintf(&result, "\n%s", card)
	}
	return result.String()
}
