package main
import "fmt"

func main(){
	// var card string ="Ace of Spades"
	// cards :=deck{"Ace of Spades",newCard()}
	// cards =append(cards,"six of spades")
	// card := "Ace of Spades"
	cards:=newDeck()
	fmt.Println(cards)
	cards.print()

	hand, remainingCards := deal(cards, 5)
	fmt.Println("hand:")
	hand.print()
	fmt.Println("Remaining:")
	remainingCards.print()
    

	fmt.Println("toString Test:",cards.toString())
	cards.saveToFile("my_cards")

	DeckFromFilecards:=newDeckFromFile("my_cards")
	fmt.Println("new Deck From File :",DeckFromFilecards)

	fmt.Println("shuffle:")
	cards.shuffle()
	cards.print()


}
// func newCard()string{
// 	return "Five of diamonds"
// }