package scoundrel

import (
	"errors"
	"fmt"
	"strings"

	"scoundrelWentRogue/internal/entities/cardgame"
)

const roomCapacity = 4

var errInvalidPosition = errors.New("card position out of range")

// Room holds the cards currently available in the room.
type Room struct {
	cards []cardgame.Card
	run   bool
}

// NewRoom creates an empty room.
func NewRoom() *Room {
	return &Room{cards: make([]cardgame.Card, 0, roomCapacity)}
}

// DrawRoom draws cards up to room capacity or until the deck is exhausted.
func (r *Room) DrawRoom(deck *cardgame.Deck) {
	for len(r.cards) < roomCapacity && deck.CardCount() > 0 {
		card, err := deck.PopFirst()
		if err != nil {
			break
		}
		r.cards = append(r.cards, card)
	}
	r.run = false
}

// RunAway returns the room cards to the deck and redraws once.
func (r *Room) RunAway(deck *cardgame.Deck) {
	if r.run {
		return
	}
	for _, card := range r.cards {
		deck.AddCard(card)
	}
	r.cards = r.cards[:0]
	r.DrawRoom(deck)
	r.run = true
}

// Select removes a card from the room and applies the effect for its suit.
// For monsters it returns remaining damage; for hearts and weapons it returns zero.
func (r *Room) Select(position int, hand *Hand, discard *DiscardPile) (int, error) {
	if position < 0 || position >= len(r.cards) {
		return 0, errInvalidPosition
	}
	selected := r.cards[position]
	r.cards = append(r.cards[:position], r.cards[position+1:]...)
	r.run = true
	value := combatValue(selected.Value)

	switch selected.Suit {
	case "diamonds":
		hand.Equip(selected, discard)
		return 0, nil
	case "heart":
		discard.AddCard(selected)
		return value, nil
	}

	weapon, hasWeapon := hand.Weapon()
	if hasWeapon {
		defeated := hand.DefeatedCards()
		if len(defeated) == 0 || value < combatValue(defeated[len(defeated)-1].Value) {
			hand.AddCard(selected)
			damage := weapon.Value - value
			if damage < 0 {
				damage = 0
			}
			return damage, nil
		}
	}

	discard.AddCard(selected)
	return -selected.Value, nil
}

// Check redraws the room when only one card remains.
func (r *Room) Check(deck *cardgame.Deck) {
	if len(r.cards) == 1 {
		r.DrawRoom(deck)
	}
}

// Cards returns a copy of the cards currently in the room.
func (r *Room) Cards() []cardgame.Card {
	return append([]cardgame.Card(nil), r.cards...)
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

func combatValue(value int) int {
	if value == 1 {
		return 14
	}
	return value
}
