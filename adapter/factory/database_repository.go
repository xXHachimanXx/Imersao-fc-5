package factory

import (
	"database/sql"

	repo "github.com/xXHachimanXx/Imersao-fc-5/adapter/repository"
	"github.com/xXHachimanXx/Imersao-fc-5/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDb(r.DB)
}
