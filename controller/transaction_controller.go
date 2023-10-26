package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Axrous/mnc/helper"
	"github.com/Axrous/mnc/model/web"
	"github.com/Axrous/mnc/service"
	"github.com/julienschmidt/httprouter"
)

type TransactionController interface{
	Payment(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type transactionController struct {
	service service.TransactionService
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
	writer.WriteHeader(http.StatusCreated)
	helper.WriteToResponseBody(writer, webResponse)
}

func NewTransactionController(service service.TransactionService) TransactionController {
	return &transactionController{
		service: service,
	}
}