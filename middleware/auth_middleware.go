package middleware

import (
	"context"
	"net/http"

	"github.com/Axrous/mnc/helper"
	"github.com/Axrous/mnc/model/web"
)

type authMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) http.Handler {
	return &authMiddleware{
		Handler: handler,
	}
}

func (middleware *authMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	//you can define auth in here like jwt

	/*
		if unauthorized you can give respon unautorized
		if valid authentication you can go next to handler like below
	*/

	path := request.URL.Path

	if path == "/api/v1/login" || path == "/api/v1/register" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	tokenString := request.Header.Get("Authorization")
	if tokenString == "" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	getData, err := helper.CompareToken(tokenString)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   err.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	ctx := context.WithValue(request.Context(), "id", getData)

	middleware.Handler.ServeHTTP(writer, request.WithContext(ctx))
}