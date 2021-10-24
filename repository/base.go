package repository

import (
	"spos/stores/models"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

type Repository interface {
	// FindStoreByDomain is query will return store with condition
	// domain is equal from request
	FindStoreByDomain(domain string) (*models.Store, error)

	// FindStoreByname is query and will return store with condition
	// domain is equal from request
	FindStoreByName(domain string) (*models.Store, error)

	// InsertStore query for create new store into database
	InsertStore(store *models.Store) (*models.Store, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
