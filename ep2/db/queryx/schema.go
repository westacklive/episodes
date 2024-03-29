// Code generated by queryx, DO NOT EDIT.

package queryx

type Schema struct {
	Card              *Table
	CardID            *BigIntColumn
	CardName          *StringColumn
	CardCost          *IntegerColumn
	CardPower         *IntegerColumn
	CardAbility       *StringColumn
	CardDefID         *StringColumn
	CardCreatedAt     *DatetimeColumn
	CardUpdatedAt     *DatetimeColumn
	DeckCard          *Table
	DeckCardID        *BigIntColumn
	DeckCardDeckID    *BigIntColumn
	DeckCardCardID    *BigIntColumn
	DeckCardCreatedAt *DatetimeColumn
	DeckCardUpdatedAt *DatetimeColumn
	Deck              *Table
	DeckID            *BigIntColumn
	DeckName          *StringColumn
	DeckCreatedAt     *DatetimeColumn
	DeckUpdatedAt     *DatetimeColumn
}

func NewSchema() *Schema {
	card := NewTable("cards")
	deckcard := NewTable("deck_cards")
	deck := NewTable("decks")

	return &Schema{
		Card:              card,
		CardID:            card.NewBigIntColumn("id"),
		CardName:          card.NewStringColumn("name"),
		CardCost:          card.NewIntegerColumn("cost"),
		CardPower:         card.NewIntegerColumn("power"),
		CardAbility:       card.NewStringColumn("ability"),
		CardDefID:         card.NewStringColumn("def_id"),
		CardCreatedAt:     card.NewDatetimeColumn("created_at"),
		CardUpdatedAt:     card.NewDatetimeColumn("updated_at"),
		DeckCard:          deckcard,
		DeckCardID:        deckcard.NewBigIntColumn("id"),
		DeckCardDeckID:    deckcard.NewBigIntColumn("deck_id"),
		DeckCardCardID:    deckcard.NewBigIntColumn("card_id"),
		DeckCardCreatedAt: deckcard.NewDatetimeColumn("created_at"),
		DeckCardUpdatedAt: deckcard.NewDatetimeColumn("updated_at"),
		Deck:              deck,
		DeckID:            deck.NewBigIntColumn("id"),
		DeckName:          deck.NewStringColumn("name"),
		DeckCreatedAt:     deck.NewDatetimeColumn("created_at"),
		DeckUpdatedAt:     deck.NewDatetimeColumn("updated_at"),
	}
}

func (s *Schema) And(clauses ...*Clause) *Clause {
	return clauses[0].And(clauses[1:]...)
}

func (s *Schema) Or(clauses ...*Clause) *Clause {
	return clauses[0].Or(clauses[1:]...)
}

func (s *Schema) ChangeCard() *CardChange {
	return &CardChange{
		ID:        BigInt{},
		Name:      String{},
		Cost:      Integer{},
		Power:     Integer{},
		Ability:   String{},
		DefID:     String{},
		CreatedAt: Datetime{},
		UpdatedAt: Datetime{},
	}
}

func (s *Schema) ChangeDeckCard() *DeckCardChange {
	return &DeckCardChange{
		ID:        BigInt{},
		DeckID:    BigInt{},
		CardID:    BigInt{},
		CreatedAt: Datetime{},
		UpdatedAt: Datetime{},
	}
}

func (s *Schema) ChangeDeck() *DeckChange {
	return &DeckChange{
		ID:        BigInt{},
		Name:      String{},
		CreatedAt: Datetime{},
		UpdatedAt: Datetime{},
	}
}
