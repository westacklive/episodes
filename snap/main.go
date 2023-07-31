package main

import (
	"encoding/json"
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

func main() {
	if err := createCards(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	client, err := db.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	c = client
}
