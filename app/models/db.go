package models

import (
	"context"
	"database/sql"
)

const DBContextKey = "db"

func MustUseDB(ctx context.Context) *DB {
	if db, ok := ctx.Value(DBContextKey).(*DB); ok {
		return db
	}
	panic("db is not set in context")
}

type DB struct {
	db        *sql.DB
	currentTx *sql.Tx
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (d *DB) Begin() error {
	if d.currentTx != nil {
		return nil
	}
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	d.currentTx = tx
	return nil
}

func (d *DB) Commit() error {
	if d.currentTx == nil {
		return nil
	}
	err := d.currentTx.Commit()
	if err != nil {
		return err
	}
	d.currentTx = nil
	return nil
}

func (d *DB) Rollback() error {
	if d.currentTx == nil {
		return nil
	}
	err := d.currentTx.Rollback()
	if err != nil {
		return err
	}
	d.currentTx = nil
	return nil
}

func (d *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if d.currentTx != nil {
		return d.currentTx.Query(query, args...)
	}
	return d.db.Query(query, args...)
}

func (d *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if d.currentTx != nil {
		return d.currentTx.Exec(query, args...)
	}
	return d.db.Exec(query, args...)
}
