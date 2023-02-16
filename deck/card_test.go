package deck

import (
	"math/rand"
	"testing"
)

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

func TestDefaultSort(t *testing.T) {
	got := New(DefaultSort)[0]
	want := Card{Rank: Ace, Suit: Spade}
	if got != want {
		t.Errorf("Expected %s as first card, got %s", want.String(), got.String())
	}
}

func TestSort(t *testing.T) {
	got := New(Sort(Less))[0]
	want := Card{Rank: Ace, Suit: Spade}
	if got != want {
		t.Errorf("Expected %s as first card, got %s", want.String(), got.String())
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	originalDeck := New()
	cards := New(Shuffle)
	if cards[0] != originalDeck[40] {
		t.Errorf("Expected the first card to be %s, got %s", originalDeck[40], cards[0])
	}
}

func TestJokersCount(t *testing.T) {
	cards := New(Jokers(2))
	jokers := getJokers(cards)

	if len(jokers) != 2 {
		t.Errorf("Wrong numbers of Jokers, got %d want 2", len(jokers))
	}
}

func TestJokersIdentity(t *testing.T) {
	cards := New(Jokers(2))
	jokers := getJokers(cards)
	m := make(map[Suit]Card)
	for _, joker := range jokers {
		m[Suit(joker.Rank)] = joker
	}

	if len(m) != len(jokers) {
		t.Errorf("Jokers should be different, got %+v and %+v", jokers[0], jokers[1])
	}
}

func getJokers(cards []Card) []Card {
	jokers := make([]Card, 0, 2)
	for _, card := range cards {
		if card.Suit == Joker {
			jokers = append(jokers, card)
		}
	}
	return jokers
}

func TestFilter(t *testing.T) {
	filterFn := func(card Card) bool {
		return card.Suit == Heart
	}
	cards := New(Filter(filterFn))
	for _, card := range cards {
		if card.Suit == Heart {
			t.Errorf("Expected all Heart cards to be filtered out, got %+v", card)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, got %d cards.", 13*4*3, len(cards))
	}
}
