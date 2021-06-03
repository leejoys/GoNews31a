package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// Конструктор объекта хранилища.
func New() *Store {
	db, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		log.Fatal(err)
	}
	return &Store{db: db}
}

func (s *Store) Posts() ([]storage.Post, error) {
	var posts []storage.Post
	return posts, nil
}

func (s *Store) AddPost(storage.Post) error {
	return nil
}
func (s *Store) UpdatePost(storage.Post) error {
	return nil
}
func (s *Store) DeletePost(storage.Post) error {
	return nil
}
