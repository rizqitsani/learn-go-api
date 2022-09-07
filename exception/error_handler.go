package exception

import (
	"net/http"

	"github.com/rizqitsani/learn-go-api/helper"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	internalServerErrorHandler(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	if err == nil {
		return false
	}

	if exception, ok := err.(NotFoundError); ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := helper.Response{
			Code:   http.StatusNotFound,
			Status: "error",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, response)

		return true
	}

	return false
}

func internalServerErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := helper.Response{
		Code:   http.StatusInternalServerError,
		Status: "error",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, response)
}
