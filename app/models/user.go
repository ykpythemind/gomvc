package models

import (
	"context"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func ListUsers(ctx context.Context, db *DB) ([]*User, error) {
	rows, err := db.Query("SELECT id, name FROM users", nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *User) FullName(ctx context.Context) (string, error) {
	fullname := fmt.Sprintf("full name: %d(%s)", u.ID, u.Name)
	return fullname, nil
}
