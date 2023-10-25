package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Axrous/mnc/model/web"
	"github.com/Axrous/mnc/service"
	"github.com/julienschmidt/httprouter"
)

type transactionController struct {
	service service.TransactionService
	router *httprouter.Router
}

func (controller *transactionController) Route() {
	controller.router.POST("/api/v1/payment", controller.Payment)
}

func (controller *transactionController) Payment(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	var transactionPaymentRequest web.TransactionPaymentRequest
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&transactionPaymentRequest)
	if err != nil {
		panic(err)
	}

	controller.service.Payment(request.Context(), transactionPaymentRequest)

	webResponse := web.WebResponse{
		Code:   201,
		Status: "CREATED",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func NewTransactionController(service service.TransactionService, router *httprouter.Router) *transactionController {
	return &transactionController{
		service: service,
		router:  router,
	}
}