package manager

import "github.com/Axrous/mnc/controller"

type ControllerManager interface {
	CustomerController() controller.CustomerController
	TransactionController() controller.TransactionController
}

type controllerManagerImpl struct {
	serviceManager ServiceManager
}

// CustomerController implements ControllerManager.
func (controllerManager *controllerManagerImpl) CustomerController() controller.CustomerController {
	return controller.NewCustomerController(controllerManager.serviceManager.CustomerService())
}

// TransactionController implements ControllerManager.
func (controllerManager *controllerManagerImpl) TransactionController() controller.TransactionController {
	return controller.NewTransactionController(controllerManager.serviceManager.TransactionService())
}

func NewControllerManager(serviceManager ServiceManager) ControllerManager {
	return &controllerManagerImpl{
		serviceManager: serviceManager,
	}
}
