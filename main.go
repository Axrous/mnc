package main

import (
	"net/http"

	"github.com/Axrous/mnc/app"
	"github.com/Axrous/mnc/exception"
	"github.com/Axrous/mnc/manager"
	"github.com/Axrous/mnc/middleware"
)

func main()  {
	
	db := app.NewDB()
	
	repositoryManager := manager.NewRepositorymanager(db)
	serviceManager := manager.NewServiceManager(repositoryManager)
	controllerManager := manager.NewControllerManager(serviceManager)
	
	router := app.NewRouter(controllerManager)
	router.PanicHandler = exception.ErrorHandler
	handler := middleware.NewLoggerMiddleware(middleware.NewAuthMiddleware(router))
	
	// handler := ex
	serve := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}