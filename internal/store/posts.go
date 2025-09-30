package store

import (
	"context"
	"database/sql"

	"github.com/go-pg/pg"
)

type Post struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (store *PostStore) Create(context context.Context, post *Post) error {
	query := `
	INSERT INTO posts (title, content, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURENING id, created_at, updated_At
	`
	err := store.db.QueryRowContext(
		context,
		query,
		post.Title,
		post.Content,
		post.UserID,
		pg.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil

}
