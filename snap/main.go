package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"snap/db"
)

var c *db.QXClient

type Download struct {
	Success struct {
		Cards []struct {
			Name      string `json:"name"`
			Cost      int    `json:"cost"`
			Power     int    `json:"power"`
			Ability   string `json:"ability"`
			Carddefid string `json:"carddefid"`
		} `json:"cards"`
	} `json:"success"`
}

func createCards() error {
	var data Download
	b, err := os.ReadFile("download.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	for _, card := range data.Success.Cards {
		if card.Carddefid != "" {
			newCard := c.ChangeCard().SetName(card.Name).SetAbility(card.Ability).
				SetCost(int32(card.Cost)).SetPower(int32(card.Power)).SetDefID(card.Carddefid)
			_, err := c.QueryCard().Create(newCard)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type Deck struct {
	Name  string
	Cards []struct {
		CardDefId string
	}
}

func createDeck(code string) (*db.Deck, error) {
	decoded, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return nil, err
	}

	var data Deck
	if err := json.Unmarshal(decoded, &data); err != nil {
		return nil, err
	}

	tx, err := c.Tx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	deck, err := tx.QueryDeck().Create(tx.ChangeDeck().SetName(data.Name))
	if err != nil {
		return nil, err
	}

	for _, cd := range data.Cards {
		card, err := tx.QueryCard().Where(tx.CardDefID.EQ(cd.CardDefId)).First()
		if err != nil {
			return nil, err
		}

		if card != nil {
			_, err := tx.QueryDeckCard().Create(tx.ChangeDeckCard().SetCardID(card.ID).SetDeckID(deck.ID))
			if err != nil {
				return nil, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return deck, nil
}

func main() {
	code := "eyJOYW1lIjoiZGlub2Z1cnlzaG9wIiwiQ2FyZHMiOlt7IkNhcmREZWZJZCI6IlN1bnNwb3QifSx7IkNhcmREZWZJZCI6IkFnZW50MTMifSx7IkNhcmREZWZJZCI6IkljZW1hbiJ9LHsiQ2FyZERlZklkIjoiS2l0dHlQcnlkZSJ9LHsiQ2FyZERlZklkIjoiQW5nZWxhIn0seyJDYXJkRGVmSWQiOiJUaGVDb2xsZWN0b3IifSx7IkNhcmREZWZJZCI6Ik15c3RpcXVlIn0seyJDYXJkRGVmSWQiOiJCaXNob3AifSx7IkNhcmREZWZJZCI6IkFnZW50Q291bHNvbiJ9LHsiQ2FyZERlZklkIjoiU2hhbmdDaGkifSx7IkNhcmREZWZJZCI6Ik5pY2tGdXJ5In0seyJDYXJkRGVmSWQiOiJEZXZpbERpbm9zYXVyIn1dfQ=="
	deck, err := createDeck(code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deck)

	deck, _ = c.QueryDeck().PreloadCards().Find(deck.ID)
	fmt.Println(deck.Cards)

	// if err := createCards(); err != nil {
	// 	log.Fatal(err)
	// }
}

func init() {
	client, err := db.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	c = client
}
