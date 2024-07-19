package middleware

import (
	"akmmp241/belajar-golang-restful-api/helper"
	"akmmp241/belajar-golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-KEY") != "RAHASIA" {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.Response{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
		}

		helper.WriteToResponseBody(writer, &webResponse)
		return
	}

	middleware.Handler.ServeHTTP(writer, request)
}
