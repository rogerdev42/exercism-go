package blackjack

const (
    SPLIT = "P"
    WIN   = "W" 
    STAND = "S"
    HIT   = "H"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
    case "two": return 2
    case "three": return 3
    case "four": return 4
    case "five": return 5
    case "six": return 6
    case "seven": return 7
    case "eight": return 8
    case "nine": return 9
    case "ten", "jack", "queen", "king": return 10
    case "ace": return 11
    default: return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerCardsSum := ParseCard(card1) + ParseCard(card2)
    dealerCardsSum := ParseCard(dealerCard)
    
	switch {
        case playerCardsSum >= 12 && playerCardsSum <= 16 && dealerCardsSum < 7: return STAND
        case playerCardsSum >= 17 && playerCardsSum <= 20: return STAND
        case playerCardsSum == 21 && dealerCardsSum < 10: return WIN
        case playerCardsSum == 21 && dealerCardsSum >= 10: return STAND
        case playerCardsSum == 22: return SPLIT
        default: return HIT
    }
}
