package main

import(
	"fmt"
	cardgame "scoundrelWentRogue/internal/entities/cardgame"
	//scoundrel "scoundrelWentRogue/internal/entities/scoundrel"
)

func main(){
	fmt.Println("Helloooooo")
	d:= cardgame.NewDeck()
	fmt.Print(d)
}
