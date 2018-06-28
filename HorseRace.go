package main
/*
Author: Evan Gardner
Version: 1.0
*/
import (
	"fmt"
	"math/rand"
	"time"
)
var deck [52]card
var playerHand [5]card
var compHand [5]card
var spadeCount int
var heartCount int
var diamondCount int
var clubCount int
var deckTop int = 0
var deckEnd int = 52
var cont string = "yes"
var roundNum int
var horses [4][6]string
const deckSize = 52
const rankSize = 13
const suitSize = 4

type card struct {
	suit string
	rank int
}

func generateDeck() {
	for i := 0; i < suitSize; i++ {
		for x := 0; x < rankSize; x++ {
			switch i {
			case 0:
				deck[i * 13 + x] = card{"spades", x + 1}
			case 1:
				deck[i * 13 + x] = card{"clubs", x + 1}
			case 2:
				deck[i * 13 + x] = card{"hearts", x + 1}
			case 3:
				deck[i * 13 + x] = card{"diamonds", x + 1}
			}
		}
	}
}

func intializeHorses() {
	horses[0][5] = "|"
	horses[1][5] = "|"
	horses[2][5] = "|"
	horses[3][5] = "|"
	for i := 0; i < 4; i++ {
		for k := 0; k < 5; k++ {
			horses[i][k] = "______"
		}
	}
}

func updateHorses() {
	for k := 0; k <= 5; k++ {
		if spadeCount == k {
			horses[0][k] = "SPADES"
			if spadeCount > 0{
			horses[0][spadeCount-1] = "_____"
			}
		}
		if clubCount == k {
			horses[1][k] = "CLUBS"
			if clubCount > 0 {
			horses[1][clubCount-1] = "_____"
			}
		}
		if heartCount == k {
			horses[2][k] = "HEARTS"
			if heartCount > 0 {
			horses[2][heartCount-1] = "_____"
			}
		}
		if diamondCount == k {
			horses[3][k] = "DIAMONDS"
			if diamondCount > 0 {
			horses[3][diamondCount-1] = "______"
			}
		}
	}
}


func printHorses() {
	for i := 0; i < 4; i++ {
		for k := 0; k < 6; k++ {
			fmt.Print(horses[i][k])
		}
		fmt.Println()
	}
}
func shuffleDeck() {
	rand.Seed(time.Now().UTC().UnixNano())
	halfDeck1, halfDeck2 := copyDeck(deck)
	for j := 0; j < rand.Intn(10000); j++ {
		random := rand.Intn(200) % 26
		random2 := rand.Intn(200) % 26
		x  := halfDeck1[random]
		halfDeck1[random] = halfDeck2[random2]
		halfDeck2[random2] = x
	}
	writeDeck(halfDeck1, halfDeck2)
}

func copyDeck(deck [52]card) (x, k [26]card) {
	for i := 0; i < 26; i++ {
		x[i] = deck[i]
	}

	for d := 0; d < 26; d++ {
		k[d] = deck[d + 26]
	}
	return x, k

}

func writeDeck(half1, half2 [26]card) {
	for i := 0; i < 26; i++ {
		deck[i] = half1[i]
		deck[i + 26] = half2[i]
	}
}
func printDeck() {
	for i := 0; i < 52; i++ {
		if i % 10 == 0 {
			fmt.Print("\n")
		}
		fmt.Print(deck[i])
	}

}

func chooseCard() (x card) {
	x = deck[deckTop]
	deckTop = (deckTop+1) % 52
	return x
}

func collectCards() {
	for i := 0; i < 5; i+=2 {
		deck[(deckEnd + i) % 52] = playerHand[i]
		deck[(deckEnd + i + 1) % 52] = compHand[i]
	}
}

func update() {
        choice := chooseCard()
	fmt.Println("we got ", choice.suit, " movin forward!")
	switch choice.suit {
	case "spades" :
	spadeCount++
	case "clubs" :
	clubCount++
	case "hearts" :
	heartCount++
	case "diamonds" :
	diamondCount++
	}
	fmt.Println("spades: ", spadeCount, " club: ", clubCount, "\nhearts: ", heartCount, "diamond: ", diamondCount)
	if diamondCount >= 5 {
	fmt.Println("DIAMONDS WIN!")
	cont = "end"
	}
	if heartCount >= 5 {
	fmt.Println("HEARTS WIN!")
	cont = "end"
	}
	if clubCount >= 5 {
	fmt.Println("CLUBS WIN!")
	cont = "end"
	}
	if spadeCount >= 5 {
	fmt.Println("SPADES WIN!")
	cont = "end"
	}
	if spadeCount > roundNum && heartCount > roundNum && clubCount > roundNum && diamondCount > roundNum {
	fmt.Println("\nWho's goin back who's goin back...")
	time.Sleep(time.Second)
	choice = chooseCard()
	fmt.Println(choice.suit, " going back!")
	switch choice.suit {
	case "spades" :
	horses[0][spadeCount] = "_____"
	spadeCount--
	case "clubs" :
	horses[1][clubCount] = "_____"
	clubCount--
	case "hearts" :
	horses[2][heartCount] = "_____"
	heartCount--
	case "diamonds" :
	horses[3][diamondCount] = "_____"
	diamondCount--
	}
	roundNum++
	fmt.Println("spades: ", spadeCount, " club: ", clubCount, "\nhearts: ", heartCount, "diamond: ", diamondCount)
}



}

func main() {

	generateDeck() //Initializes the cards in the deck in order
	intializeHorses()
	updateHorses()
	fmt.Println("WELCOME TO HORSE RACE")
	fmt.Println("\nPress enter to draw another card, type 'no' if you want to stop. When a horse\ngets 5 cards they have won! Place your bets! Press enter begin")
	fmt.Scanln(&cont)
	shuffleDeck() //shuffles the cards
	for cont != "no" {
		update()
		updateHorses()
		printHorses()
		if cont == "end" {
		break
		}
		fmt.Println("Continue?")
		fmt.Scanln(&cont)
	}
	for {
	}
}

