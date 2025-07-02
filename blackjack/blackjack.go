package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "ten":
		return 10
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
	default:
		return 0 // Invalid card
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerSum := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
	// If you have a pair of aces you must always split them.
	case card1 == "ace" && card2 == "ace":
		return "P"

	// If you have a Blackjack (21), and the dealer does not have an ace, face card, or ten, you automatically win.
	case playerSum == 21 && dealerValue != 11 && dealerValue != 10:
		return "W"
	// If you have a Blackjack (21), and the dealer has ace, face card, or ten, stand and wait for reveal
	case playerSum == 21:
		return "S"

	// If your cards sum up to [17, 20] you should always stand.
	case playerSum >= 17 && playerSum <= 20:
		return "S"

	// If your cards sum up to [12, 16] you should always stand unless the dealer has a 7 or higher, then hit.
	case playerSum >= 12 && playerSum <= 16 && dealerValue >= 7:
		return "H"
	case playerSum >= 12 && playerSum <= 16:
		return "S"

	// If your cards sum up to 11 or lower you should always hit.
	case playerSum <= 11:
		return "H"

	// Default to stand (should not be reached)
	default:
		return "S"
	}
}
