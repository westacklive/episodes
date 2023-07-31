database "db" {
  adapter = "postgresql"

  config "development" {
    url = "postgres://postgres:postgres@localhost:5432/snap_development?sslmode=disable"
  }

  generator "client-golang" {}

  model "Card" {
    has_many "deck_cards" {}
    has_many "deck" {
      through = "deck_cards"
    }

    column "name" {
      type = string
      null = false
    }
    column "cost" {
      type = integer
      null = false
    }
    column "power" {
      type = integer
      null = false
    }
    column "ability" {
      type = string
      null = false
    }
    column "def_id" {
      type = string
      null = false
    }

    index {
      columns = ["def_id"]
      unique  = true
    }
  }

  model "DeckCard" { # deck_cards
    belongs_to "deck" {}
    belongs_to "card" {}
  }

  model "Deck" {
    has_many "deck_cards" {}
    has_many "cards" {
      through = "deck_cards"
    }

    column "name" {
      type = string
    }
  }
}
