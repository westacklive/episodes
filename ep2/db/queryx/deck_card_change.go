// Code generated by queryx, DO NOT EDIT.

package queryx

type DeckCardChange struct {
	ID        BigInt
	DeckID    BigInt
	CardID    BigInt
	CreatedAt Datetime
	UpdatedAt Datetime
}

func (c *DeckCardChange) Changes() (columns []string, values []interface{}) {
	if c == nil {
		return columns, values
	}
	if c.ID.Set {
		columns = append(columns, "id")
		values = append(values, c.ID)
	}
	if c.DeckID.Set {
		columns = append(columns, "deck_id")
		values = append(values, c.DeckID)
	}
	if c.CardID.Set {
		columns = append(columns, "card_id")
		values = append(values, c.CardID)
	}
	if c.CreatedAt.Set {
		columns = append(columns, "created_at")
		values = append(values, c.CreatedAt)
	}
	if c.UpdatedAt.Set {
		columns = append(columns, "updated_at")
		values = append(values, c.UpdatedAt)
	}
	return columns, values
}

func (c *DeckCardChange) SetID(id int64) *DeckCardChange {
	c.ID = NewBigInt(id)
	c.ID.Set = true
	return c
}

func (c *DeckCardChange) SetDeckID(deckID int64) *DeckCardChange {
	c.DeckID = NewBigInt(deckID)
	c.DeckID.Set = true
	return c
}

func (c *DeckCardChange) SetNullableDeckID(deckID *int64) *DeckCardChange {
	c.DeckID = NewNullableBigInt(deckID)
	c.DeckID.Set = true
	return c
}

func (c *DeckCardChange) SetCardID(cardID int64) *DeckCardChange {
	c.CardID = NewBigInt(cardID)
	c.CardID.Set = true
	return c
}

func (c *DeckCardChange) SetNullableCardID(cardID *int64) *DeckCardChange {
	c.CardID = NewNullableBigInt(cardID)
	c.CardID.Set = true
	return c
}

func (c *DeckCardChange) SetCreatedAt(createdAt string) *DeckCardChange {
	c.CreatedAt = NewDatetime(createdAt)
	c.CreatedAt.Set = true
	return c
}

func (c *DeckCardChange) SetUpdatedAt(updatedAt string) *DeckCardChange {
	c.UpdatedAt = NewDatetime(updatedAt)
	c.UpdatedAt.Set = true
	return c
}
