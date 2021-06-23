package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Хранилище данных.
type Store struct {
	c  *mongo.Client
	db *mongo.Database
}

// Конструктор объекта хранилища.
func New(name string, connstr string) (*Store, error) {
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
	return &Store{c: client, db: client.Database(name)}, nil
}

func (s *Store) Posts() ([]storage.Post, error) {

	coll := s.db.Collection("posts")
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

func (s *Store) AddPost(p storage.Post) error {
	p.CreatedAt = time.Now().Unix()
	coll := s.db.Collection("posts")
	_, err := coll.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdatePost(p storage.Post) error {
	coll := s.db.Collection("posts")
	filter := bson.E{Key: "id", Value: p.ID}
	update := bson.E{Key: "$set", Value: bson.D{
		{Key: "Title", Value: p.Title},
		{Key: "Content", Value: p.Content},
		{Key: "AuthorID", Value: p.AuthorID},
		{Key: "PublishedAt", Value: p.PublishedAt}}}
	_, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeletePost(p storage.Post) error {
	coll := s.db.Collection("posts")
	filter := bson.D{{Key: "id", Value: p.ID}}
	_, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
