package postgres

import (
	"GoNews/pkg/storage"
	"context"

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
	return &Store{db: db}, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(),
		`SELECT 
	posts.id, 
	posts.title, 
	posts.content, 
	posts.author_id,
	author.name, 
	posts.create_at 
	FROM posts
	JOIN authors
	AT authors.name=posts.author_id;`)

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
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, rows.Err()
}

func (s *Store) AddPost(p storage.Post) error {
	err := s.db.QueryRow(SELECT)

	err = s.db.Exec("INSERT INTO posts ", p.Title, p.Content, p.AuthorID, p.CreatedAt)
	return nil
}
func (s *Store) UpdatePost(storage.Post) error {
	err := s.db.Exec("DELETE FROM posts WHERE ", p.Title, p.Content, p.AuthorID, p.CreatedAt)
	return nil
}
func (s *Store) DeletePost(storage.Post) error {
	err := s.db.Exec("DELETE FROM posts WHERE ", p.Title, p.Content, p.AuthorID, p.CreatedAt)
	return nil
}
