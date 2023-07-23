package main

import (
	"ep1/db"
	"fmt"
	"log"
)

func main() {
	c, err := db.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// create post
	newPost := c.ChangePost().SetTitle("post title").SetIsPublished(true)
	post, _ := c.QueryPost().Create(newPost)
	fmt.Println(post)
}
