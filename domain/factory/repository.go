package factory

import "github.com/xXHachimanXx/Imersao-fc-5/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
