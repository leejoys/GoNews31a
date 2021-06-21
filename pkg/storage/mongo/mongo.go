package mongo

import (
	"GoNews/pkg/storage"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Хранилище данных.
type Store struct {
	c          *mongo.Client
	db         string
	collection string
}

// Конструктор объекта хранилища.
func New(connstr string) (*Store, error) {
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(connstr))
	if err != nil {
		return nil, err
	}
	// проверка связи с БД
	err = client.Ping(context.Background(), nil)
	if err != nil {
		client.Disconnect(context.Background())
		return nil, err
	}
	return &Store{c: client, db: "data",
		collection: "posts"}, nil
}

func (s *Store) Posts() ([]storage.Post, error) {

	coll := s.c.Database(s.db).Collection(s.collection)
	ctx := context.Background()
	filter := bson.D{}
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var posts []storage.Post
	for cur.Next(ctx) {
		var p storage.Post
		err = cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
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
