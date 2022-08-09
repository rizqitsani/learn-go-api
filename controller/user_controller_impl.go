package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rizqitsani/learn-go-api/dto"
	"github.com/rizqitsani/learn-go-api/helper"
	"github.com/rizqitsani/learn-go-api/service"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createUserRequest := dto.CreateUserDto{}
	helper.ReadFromRequestBody(request, &createUserRequest)

	user := controller.userService.Create(request.Context(), createUserRequest)
	response := helper.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, response)
}
