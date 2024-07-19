package exception

import (
	"akmmp241/belajar-golang-restful-api/helper"
	"akmmp241/belajar-golang-restful-api/model/web"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if validationErrors(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)

	webResponse := web.Response{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Data:   exception.Error(),
	}

	helper.WriteToResponseBody(writer, &webResponse)
	return true
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if !ok {
		return false
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)

	webResponse := web.Response{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Data:   exception.Error,
	}

	helper.WriteToResponseBody(writer, &webResponse)
	return true
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.Response{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Data:   err,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}
