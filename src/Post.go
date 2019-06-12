package main

import (
	"fmt"
	"time"
)

type Post struct {
	ID             int
	Title          string
	Content        string
	ContentRecords []ContentRecord
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Excerpt        string
	ForumAlias     string
	Gender         string
	School         string
	Comments       map[string]Comment
}

type ContentRecord struct {
	Content   string
	UpdatedAt time.Time
}

func (p *Post) addRecord(post *Post) {
	// 避免重複的紀錄
	for _, contentRecord := range p.ContentRecords {
		if post.UpdatedAt == contentRecord.UpdatedAt {
			return
		}
	}

	// 紀錄內容記錄
	p.ContentRecords = append(p.ContentRecords, ContentRecord{
		Content:   post.Content,
		UpdatedAt: post.UpdatedAt,
	})
}

type Comment struct {
	ID        string
	Anonymous bool
	Floor     int
	Content   string
	Records   []Comment
	LikeCount int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Comment) addRecord(comment Comment) {
	// 避免重複的紀錄
	for _, record := range c.Records {
		if comment.UpdatedAt == record.UpdatedAt {
			return
		}
	}

	// 紀錄內容記錄
	c.Records = append(c.Records, comment)
}

func (p *Post) loadPost() {
	dcard.getPost(p)

	p.addRecord(p)
}

func (p *Post) setComment(comment Comment) {
	floorString := fmt.Sprintf("%d", comment.Floor)
	fmt.Println(floorString)

	_, ok := p.Comments[floorString]
	if ok {
		comment.Records = p.Comments[floorString].Records
		comment.addRecord(comment)
		p.Comments[floorString] = comment
	} else {
		comment.addRecord(comment)
		p.Comments[floorString] = comment
	}
}

func (p *Post) loadComments() {
	if len(p.Comments) == 0 {
		p.Comments = make(map[string]Comment)
	}

	dcard.getPostComment(p)
}
