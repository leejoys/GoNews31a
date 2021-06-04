package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// Конструктор объекта хранилища.
func New() *Store {
	pwd := os.Getenv("pgpass")
	connstr := "postgres://postgres:" + pwd + "@server.domain/database"
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
