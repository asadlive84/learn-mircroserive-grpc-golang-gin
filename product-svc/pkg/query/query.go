package query

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type QueryInit struct {
	Db *sqlx.DB
}

var NotFound = errors.New("not found")

type Product struct {
	ID          string       `db:"id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	IsActive    bool         `db:"is_active"`
	Created     time.Time    `db:"created"`
	Updated     sql.NullTime `db:"updated"`
	CreatedBy   string       `db:"created_by"`
}

type Query interface {
	InsertProduct(u Product) error
	GetProductByID(email string) (*Product, error)
}
