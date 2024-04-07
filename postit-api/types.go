package main

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID        string      `json:"id"`
	Username  string      `json:"username"`
	Password  string      `json:"password"`
	CreatedAt time.Time   `json:"created_at"`
	Avatar    null.String `json:"avatar"`
	Bio       null.String `json:"bio"`
}

type Board struct {
	ID          int64       `json:"id,string"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Thumbnail   null.String `json:"thumbnail"`
	CreatedAt   time.Time   `json:"created_at"`
	CreatorId   string      `json:"creatorId"`
}
