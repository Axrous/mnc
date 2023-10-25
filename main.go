package main

import (
	"net/http"

	"github.com/Axrous/mnc/app"
	"github.com/Axrous/mnc/controller"
	"github.com/Axrous/mnc/exception"
	"github.com/Axrous/mnc/middleware"
	"github.com/Axrous/mnc/repository"
	"github.com/Axrous/mnc/service"
	"github.com/julienschmidt/httprouter"
)

func main()  {
	
	db := app.NewDB()
	router := httprouter.New()


	customerRepository := repository.NewCustomerRepository(db)
	whitelistRepository := repository.NewWhiteListRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	customerService := service.NewCustomerService(customerRepository, whitelistRepository)
	transactionService := service.NewTransactionService(transactionRepository, whitelistRepository)

	router.PanicHandler = exception.ErrorHandler
	controller.NewCustomerController(customerService, router).Route()
	controller.NewTransactionController(transactionService, router).Route()
	
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