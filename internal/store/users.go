package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updted_at"`
}

type UserStore struct {
	db *sql.DB
}

func (store *UserStore) Create(context context.Context, user *User) error {

	query := `
	INSERT INTO users (username, email, password)
	VALUES ($1, $2, $3) RETURENING id, created_at
	`
	err := store.db.QueryRowContext(
		context,
		query,
		user.UserName,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}
