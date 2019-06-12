package main

import (
	"fmt"
)

var dcard = Dcard{"www.dcard.cc"}

func main() {
	posts := []Post{}
	dcard.getPopularPosts("food", 5, &posts)

	posts[3].loadPost()
	posts[3].loadComments()
	fmt.Println(posts[3].Comments)


	mongo := Mongo{}
	mongo.login()
	mongo.savePost(posts[3])

	// for _, post := range posts {
	// fmt.Println(post)
	// }
}
