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
