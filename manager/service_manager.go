package manager

import (
	"github.com/Axrous/mnc/service"
)

type ServiceManager interface {
	CustomerService() service.CustomerService
	TransactionService() service.TransactionService
}

type serviceManagerImpl struct {
	repoManager RepositoryManager
}

// CustomerService implements ServiceManager.
func (serviceManager *serviceManagerImpl) CustomerService() service.CustomerService {
	return service.NewCustomerService(serviceManager.repoManager.CustomerRepository(), serviceManager.repoManager.WhitelistRepository())
}

// TransactionService implements ServiceManager.
func (serviceManager *serviceManagerImpl) TransactionService() service.TransactionService {
	return service.NewTransactionService(serviceManager.repoManager.TransactionRepository(), serviceManager.repoManager.WhitelistRepository())
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	return &serviceManagerImpl{
		repoManager: repoManager,
	}
}
