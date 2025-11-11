package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int

	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		value = 10
	case "jack":
		value = 10
	case "queen":
		value = 10
	case "king":
		value = 10
	case "joker":
		value = 0
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	if handValue == 21 && dealerValue >= 10 {
		return "S"
	}

	if handValue <= 20 && handValue >= 17 {
		return "S"
	}

	if handValue <= 16 && handValue >= 12 {
		if dealerValue >= 7 {
			return "H"
		}

		return "S"
	}

	if handValue <= 11 {
		return "H"
	}

	if handValue == 21 && dealerCard != "ace" || handValue == 21 && dealerValue != 10 {
		return "W"
	}

	return "S"
}
