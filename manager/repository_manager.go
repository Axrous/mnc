package manager

import (
	"github.com/Axrous/mnc/repository"
	"github.com/sonyarouje/simdb"
)

type RepositoryManager interface {
	CustomerRepository() repository.CustomerRepository
	WhitelistRepository() repository.WhitelistRepository
	TransactionRepository() repository.TransactionRepository
}

type repositoryManagerImpl struct {
	db *simdb.Driver
}

// CustomerRepository implements RepositoryManager.
func (repoManager *repositoryManagerImpl) CustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(repoManager.db)
}

// TransactionRepository implements RepositoryManager.
func (repoManager *repositoryManagerImpl) TransactionRepository() repository.TransactionRepository {
	return repository.NewTransactionRepository(repoManager.db)
}

// WhitelistRepository implements RepositoryManager.
func (repoManager *repositoryManagerImpl) WhitelistRepository() repository.WhitelistRepository {
	return repository.NewWhiteListRepository(repoManager.db)
}

func NewRepositorymanager(db *simdb.Driver) RepositoryManager {
	return &repositoryManagerImpl{
		db: db,
	}
}
