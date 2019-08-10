package card

// Card represents a single flashcard. A card has two sides, front and back.
// A card may be reversible. If a card is reversible, either side may be shown
// as the prompt. Otherwise, the front side is the prompt, and the back is the
// solution.
type Card struct {
	Front      string
	Back       string
	Reversible bool
}

// New returns a new card.
func New(f, b string, r bool) *Card {
	return &Card{
		Front:      f,
		Back:       b,
		Reversible: r,
	}
}
