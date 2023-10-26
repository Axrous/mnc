package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Axrous/mnc/helper"
	"github.com/Axrous/mnc/model/web"
	"github.com/Axrous/mnc/service"
	"github.com/julienschmidt/httprouter"
)

type CustomerController interface{
	CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetCustById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type customerControllerImpl struct {
	service service.CustomerService
}


func (controller *customerControllerImpl) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	var customerCreateRequest web.CustomerCreateRequest
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&customerCreateRequest)
	if err != nil {
		panic(err)
	}

	controller.service.Save(request.Context(), customerCreateRequest)

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

func (controller *customerControllerImpl) GetCustById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")

	customer := controller.service.FindById(request.Context(), userId)

	webResponse := web.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data: customer,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *customerControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	var customerLoginRequest web.CustomerLoginRequest
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&customerLoginRequest)
	if err != nil {
		panic(err)
	}

	token := controller.service.Login(request.Context(), customerLoginRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data: token,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *customerControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	
	controller.service.Logout(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data: "Logout",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func NewCustomerController(service service.CustomerService) CustomerController {
	return &customerControllerImpl{
		service: service,
	}
}