package main

import (
	"fmt"
)

var dcard = Dcard{"www.dcard.cc"}

func main() {
	posts := []Post{}
	dcard.getPopularPosts("food", 5, &posts)

	for _, post := range posts {
		fmt.Println(post)
	}
}
