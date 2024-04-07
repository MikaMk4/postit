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
