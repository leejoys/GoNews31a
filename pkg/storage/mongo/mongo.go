package mongo

import (
	"GoNews/pkg/storage"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Хранилище данных.
type Store struct {
	db *mongo.Client
}

// Конструктор объекта хранилища.
func New(connstr string) (*Store, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connstr))
	if err != nil {
		return nil, err
	}
	return &Store{db: client}, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	//err := s.db.Database() // <--
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
