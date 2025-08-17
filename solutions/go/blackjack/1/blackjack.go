package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var cardValue int
	switch card {
        case "ace":
        	cardValue = 11
        case "two":
        	cardValue = 2
        case "three":
        	cardValue = 3
        case "four":
        	cardValue = 4
        case "five":
        	cardValue = 5
        case "six":
        	cardValue = 6
        case "seven":
        	cardValue = 7
        case "eight":
        	cardValue = 8
        case "nine":
        	cardValue = 9
        case "ten", "jack", "queen", "king":
        	cardValue = 10
        default:
        	cardValue = 0
    }
    return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    const blackJack int = 21
	var option string = ""
    valueCard1 := ParseCard(card1)
    valueCard2 := ParseCard(card2)
    sumOfCards := valueCard1 + valueCard2
    switch {
        case card1 == card2 && card1 == "ace":
        	option = "P"
        case sumOfCards >= blackJack:
        	switch dealerCard {
                case "ace", "jack", "queen", "king", "ten":
                	option = "S"
                default:
                	option = "W"
            }
        case sumOfCards >= 17 && sumOfCards <= 20:
        	option = "S"
        case sumOfCards >= 12 && sumOfCards <= 16:
        	option = "S"
        	if ParseCard(dealerCard) >= 7 {
                option = "H"
            }
        case sumOfCards <= 11:
        	option = "H"
    }
    return option
}
