package app

import (
	"github.com/Axrous/mnc/manager"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(controllerManager manager.ControllerManager) *httprouter.Router{

	router := httprouter.New()

	//Customer Controller
	router.POST("/api/v1/register", controllerManager.CustomerController().CreateUser)
	router.GET("/api/v1/customers/:id", controllerManager.CustomerController().GetCustById)
	router.POST("/api/v1/login", controllerManager.CustomerController().Login)
	router.POST("/api/v1/logout", controllerManager.CustomerController().Logout)

	//Transaction Controller
	router.POST("/api/v1/payment", controllerManager.TransactionController().Payment)


	return router
}