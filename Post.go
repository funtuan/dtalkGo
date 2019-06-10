package main

import (
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
	Comments       []Comment
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
	floor := comment.Floor
	if len(p.Comments) >= floor {
		comment.Records = p.Comments[floor-1].Records
		p.Comments[floor-1] = comment
		p.Comments[floor-1].addRecord(comment)
	} else {
		comment.addRecord(comment)
		p.Comments = append(p.Comments, comment)
	}
}

func (p *Post) loadComments() {
	dcard.getPostComment(p)
}
