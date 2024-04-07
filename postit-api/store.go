package main

import (
	"database/sql"

	"github.com/google/uuid"
)

type Store interface {
	// Posts
	GetPosts() ([]Post, error)
	CreatePost(post *Post) (*Post, error)

	// Users
	GetUsers() ([]User, error)
	GetUserById(id string) (*User, error)
	GetUserByName(username string) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) GetPosts() ([]Post, error) {
	rows, err := s.db.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *Storage) CreatePost(post *Post) (*Post, error) {
	rows, err := s.db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", post.Title, post.Content)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	post.ID = id
	return post, nil
}

func (s *Storage) GetUsers() ([]User, error) {
	rows, err := s.db.Query("SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *Storage) GetUserById(id string) (*User, error) {
	var user User
	err := s.db.QueryRow("SELECT id, username, avatar, bio, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Avatar, &user.Bio, &user.Password)
	return &user, err
}

func (s *Storage) GetUserByName(username string) (*User, error) {
	var user User
	err := s.db.QueryRow("SELECT password, id FROM users WHERE username = ?", username).Scan(&user.Password, &user.ID)
	return &user, err
}

func (s *Storage) CreateUser(user *User) (*User, error) {
	uuid := uuid.New()
	id := uuid.String()

	_, err := s.db.Exec("INSERT INTO users (username, password, id) VALUES (?, ?, ?)", user.Username, user.Password, id)

	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}

func (s *Storage) UpdateUser(user *User) (*User, error) {
	_, err := s.db.Exec("UPDATE users SET username = ?, password = ?, avatar = ?, bio = ? WHERE id = ?", user.Username, user.Password, user.Avatar, user.Bio, user.ID)
	return user, err
}
