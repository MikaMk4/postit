package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MySqlStorage struct {
	db *sql.DB
}

func NewMySqlStorage(cfg mysql.Config) *MySqlStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	for err = db.Ping(); err != nil; {
		log.Println("Waiting for MySQL to be ready... (" + err.Error() + ")")
		time.Sleep(2 * time.Second)
	}

	log.Println("Connected to MySQL")

	return &MySqlStorage{
		db: db,
	}
}

func (s *MySqlStorage) Init() (*sql.DB, error) {
	// create tables
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createPostsTable(); err != nil {
		return nil, err
	}

	if err := s.createBoardsTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MySqlStorage) createUsersTable() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		avatar TEXT,
		bio TEXT
	)`)
	if err != nil {
		return err
	}

	return nil
}

func (s *MySqlStorage) createPostsTable() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS posts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		author_id VARCHAR(255) NOT NULL,

		FOREIGN KEY (author_id) REFERENCES users (id)
	)`)
	if err != nil {
		return err
	}

	return nil
}

func (s *MySqlStorage) createBoardsTable() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS boards (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		thumbnail TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		creator_id VARCHAR(255) NOT NULL,

		FOREIGN KEY (creator_id) REFERENCES users (id)
	)`)
	if err != nil {
		return err
	}

	return nil
}
