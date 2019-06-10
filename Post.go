package main

type Post struct {
	ID         int
	Title      string
	Excerpt    string
	ForumAlias string
	Gender     string
	School     string
}

func (*Post) loadComment() {

}
