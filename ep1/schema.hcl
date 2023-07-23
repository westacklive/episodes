database "db" {
  adapter = "postgresql"

  config "development" {
    url = "postgres://postgres:postgres@localhost:5432/ep1_development?sslmode=disable"
  }

  generator "client-golang" {}

  model "Post" {
    column "title" {
      type = string
    }
    column "content" {
      type = text
    }
    column "is_published" {
      type = boolean
    }
  }
}
