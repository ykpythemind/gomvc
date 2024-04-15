package models

import (
	"context"
	"database/sql"

	"golang.org/x/xerrors"
)

const DBContextKey = "db"

func MustUseDB(ctx context.Context) *DB {
	if db, ok := ctx.Value(DBContextKey).(*DB); ok {
		return db
	}
	panic("db is not set in context")
}

// DB is a wrapper of sql.DB. railsっぽいtransactionを提供する
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
	if d.currentTx != nil { // FIXME: this logic does not support nested transactions
		return nil
	}
	tx, err := d.db.Begin()
	if err != nil {
		return xerrors.Errorf("failed to begin transaction: %w", err)
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
		return xerrors.Errorf("failed to commit transaction: %w", err)
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
		return xerrors.Errorf("failed to rollback: %w", err)
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

func (d *DB) MustExec(query string, args ...interface{}) sql.Result {
	if d.currentTx != nil {
		result, err := d.currentTx.Exec(query, args...)
		if err != nil {
			panic(err)
		}
		return result
	}
	result, err := d.db.Exec(query, args...)
	if err != nil {
		panic(err)
	}
	return result
}
