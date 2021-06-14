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
	posts.created_at 
	FROM posts
	JOIN authors
	AT authors.id=posts.author_id;`)

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
	_, err := s.db.Exec(context.Background(), `
	INSERT INTO posts (title, content, author_id, created_at) 
	VALUES ($1,$2,$3);`, p.Title, p.Content, p.AuthorID, time.Now().Unix())
	return err
}

//UpdatePost - обновляет по id значения title, content и author_id
func (s *Store) UpdatePost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	UPDATE posts 
	SET title=$2, content=$3, author_id=$4
	WHERE id=$1;`, p.ID, p.Title, p.Content, p.AuthorID)
	return err
}

//DeletePost - удаляет пост по id
func (s *Store) DeletePost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	DELETE FROM posts 
	WHERE id=$1;`, p.ID)
	return err
}
