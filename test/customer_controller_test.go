package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Axrous/mnc/app"
	"github.com/Axrous/mnc/exception"
	"github.com/Axrous/mnc/helper"
	"github.com/Axrous/mnc/manager"
	"github.com/Axrous/mnc/middleware"
	"github.com/sonyarouje/simdb"
	"github.com/stretchr/testify/assert"
)

func setUpDB() *simdb.Driver {
	db, err := simdb.New("database-test")
	helper.PanicIfError(err)

	return db
}

func setRouter(db *simdb.Driver) http.Handler {
	
	repositoryManager := manager.NewRepositorymanager(db)
	serviceManager := manager.NewServiceManager(repositoryManager)
	controllerManager := manager.NewControllerManager(serviceManager)

	//router
	router := app.NewRouter(controllerManager)
	router.PanicHandler = exception.ErrorHandler

	return middleware.NewLoggerMiddleware(middleware.NewAuthMiddleware(router))
}

func TestCustomerLoginSuccess(t *testing.T)  {
	db := setUpDB()
	router := setRouter(db)

	requestBody := strings.NewReader(`{"username": "argasm", "password": "123"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response :=recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestCustomerLoginFailed(t *testing.T)  {
	db := setUpDB()
	router := setRouter(db)

	//wrong password
	requestBody := strings.NewReader(`{"username": "argasm", "password": "1234"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response :=recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}

func TestCustomerLogoutSuccess(t *testing.T)  {
	db := setUpDB()
	router := setRouter(db)

	//login first
	requestBody := strings.NewReader(`{"username": "argasm", "password": "123"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response :=recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	//logout then
	request = httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/logout", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", responseBody["data"].(string))

	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response =recorder.Result()
	
	assert.Equal(t, 200, response.StatusCode)

	body, _ = io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestCustomerLogoutFailed(t *testing.T)  {
	db := setUpDB()
	router := setRouter(db)


	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/logout", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response :=recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}

func TestPaymentSuccess(t *testing.T)  {
	db := setUpDB()
	router := setRouter(db)

	//login first
	requestBody := strings.NewReader(`{"username": "argasm", "password": "123"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response :=recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	//pay then
	requestBody = strings.NewReader(`{"merchant_id": "1", "amount": 5000}`)
	request = httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/payment", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", responseBody["data"].(string))

	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response =recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ = io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "CREATED", responseBody["status"])
}