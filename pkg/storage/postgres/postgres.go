package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// Конструктор объекта хранилища.
func New(connstr string) (*Store, error) {

	db, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}
	err = db.Ping(context.Background())
	if err != nil {
		db.Close()
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(),
		`SELECT 
	posts.id, 
	posts.title, 
	posts.content, 
	posts.author_id,
	authors.name, 
	posts.created_at, 
	posts.published_at
	FROM posts
	JOIN authors
	ON authors.id=posts.author_id;`)

	if err != nil {
		return nil, err
	}

	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.AuthorID,
			&p.AuthorName,
			&p.CreatedAt,
			&p.PublishedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, rows.Err()
}

func (s *Store) AddPost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	INSERT INTO posts (title, content, author_id, created_at, published_at) 
	VALUES ($1,$2,$3,$4,$5);`, p.Title, p.Content, p.AuthorID, time.Now().Unix(), time.Now().Unix())
	return err
}

//UpdatePost - обновляет по id значения title, content и author_id
func (s *Store) UpdatePost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	UPDATE posts 
	SET title=$2, content=$3, author_id=$4, published_at=$5
	WHERE id=$1;`, p.ID, p.Title, p.Content, p.AuthorID, p.PublishedAt)
	return err
}

//DeletePost - удаляет пост по id
func (s *Store) DeletePost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	DELETE FROM posts 
	WHERE id=$1;`, p.ID)
	return err
}
