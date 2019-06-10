package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type Dcard struct {
	dns string
}

func (d *Dcard) getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (d *Dcard) getPopularPosts(forumsName string, limit int, posts *[]Post) {
	d.getJson("https://"+d.dns+"/_api/forums/"+forumsName+"/posts?popular=true&limit="+fmt.Sprintf("%d", limit), &posts)
}

func (d *Dcard) getPost(post *Post) {
	d.getJson("https://"+d.dns+"/_api/posts/"+fmt.Sprintf("%d", post.ID), &post)
}

func (d *Dcard) getPostComment(post *Post) {
	floor := 0
	comments := []Comment{}

	d.getJson("https://"+d.dns+"/_api/posts/"+fmt.Sprintf("%d", post.ID)+"/comments?after="+fmt.Sprintf("%d", floor), &comments)
	for len(comments) > 0 {
		for _, comment := range comments {
			floor = comment.Floor
			post.setComment(comment)
		}
		d.getJson("https://"+d.dns+"/_api/posts/"+fmt.Sprintf("%d", post.ID)+"/comments?after="+fmt.Sprintf("%d", floor), &comments)
	}
}
