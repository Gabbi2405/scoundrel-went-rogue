package cardgame

import (
    "fmt"
    "strconv"
)

type Card struct{
    Value int
    Suit string
}

// String returns a readable representation of the card.
func (c Card) String() string{
    var name string
    switch c.Value {
    case 11: 
        name="Jack"
    case 12:
        name="Queen"
    case 13:
        name="King"
    case 1:
        name="Ace"
    case 0:
        name="joker"
    default:
        name=strconv.Itoa(c.Value)
    }
    return fmt.Sprintf("%s of %s", name, c.Suit)
}

