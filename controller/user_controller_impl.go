package controller

import (
	"net/http"
	"strconv"

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

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	users := controller.userService.FindAll(request.Context())
	response := helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   users,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	user := controller.userService.FindById(request.Context(), id)
	response := helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateUserRequest := dto.UpdateUserDto{}
	helper.ReadFromRequestBody(request, &updateUserRequest)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	updateUserRequest.Id = id

	user := controller.userService.Update(request.Context(), updateUserRequest)
	response := helper.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	controller.userService.Delete(request.Context(), id)
	response := helper.Response{
		Code:   http.StatusOK,
		Status: "success",
	}

	helper.WriteToResponseBody(writer, response)
}
