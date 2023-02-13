package deck

import "testing"

func TestString(t *testing.T) {
	testcases := []struct {
		in   Card
		want string
	}{
		{in: Card{Suit: Heart, Rank: Ace}, want: "Ace of Hearts"},
		{in: Card{Suit: Spade, Rank: Two}, want: "Two of Spades"},
		{in: Card{Suit: Diamond, Rank: Nine}, want: "Nine of Diamonds"},
		{in: Card{Suit: Club, Rank: Jack}, want: "Jack of Clubs"},
		{in: Card{Suit: Joker}, want: "Joker"},
	}
	for _, tc := range testcases {
		got := tc.in.String()
		if got != tc.want {
			t.Errorf("String: %q, want %q", got, tc.want)
		}
	}
}

func TestNew(t *testing.T) {
	got := len(New())
	want := 13 * 4
	if got != want {
		t.Errorf("Wrong numbers of cards in a new deck, got %d want %d", got, want)
	}
}
