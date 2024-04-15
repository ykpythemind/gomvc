package models

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/xerrors"
)

type User struct {
	ID   int
	Name string
}

func ListUsers(ctx context.Context, db *DB) ([]*User, error) {
	rows, err := db.Query("SELECT id, name FROM users", nil)
	if err != nil {
		return nil, xerrors.Errorf("failed to select user : %w", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, xerrors.Errorf("failed to scan user : %w", err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *User) FullName(ctx context.Context) (string, error) {
	fullname := fmt.Sprintf("full name: %d(%s)", u.ID, u.Name)
	return fullname, nil
}

// Create creates a new user. It ignore the ID field because it is auto-incremented.
func (u *User) Create(ctx context.Context, db *DB) error {
	now := time.Now()
	_, err := db.Exec("INSERT INTO users (name, created_at, updated_at) VALUES (?, ?, ?)", u.Name, now, now)
	if err != nil {
		return xerrors.Errorf("failed to insert user : %w", err)
	}
	return nil
}

// Validate validates the user model. 更新系のメソッドを呼ぶ前に明示的に呼び出す
func (u *User) Validate() error {
	if u.Name == "" {
		return xerrors.New("name is required")
	}
	return nil
}
