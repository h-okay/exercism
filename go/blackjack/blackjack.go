package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	val1 := ParseCard(card1)
	val2 := ParseCard(card2)
	dealerVal := ParseCard(dealerCard)

	// Case 1: If you have a pair of aces you must always split them.
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	// Case 2: If you have a Blackjack and the dealer does not have an ace, a figure, or a ten then you automatically win.
	if (val1+val2 == 21) && (dealerVal != 11 && dealerVal != 10) {
		return "W"
	}

	// Case 3: If your cards sum up to a value within the range [17, 20] you should always stand.
	if val1+val2 >= 17 && val1+val2 <= 20 {
		return "S"
	}

	// Case 4: If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
	if val1+val2 >= 12 && val1+val2 <= 16 {
		if dealerVal >= 7 {
			return "H"
		}
		return "S"
	}

	// Case 5: If your cards sum up to 11 or lower you should always hit.
	if val1+val2 <= 11 {
		return "H"
	}

	return "S" // Stand otherwise
}
