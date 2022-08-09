package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rizqitsani/learn-go-api/config"
	"github.com/rizqitsani/learn-go-api/controller"
	"github.com/rizqitsani/learn-go-api/repository"
	"github.com/rizqitsani/learn-go-api/service"
)

type AppController struct {
	userController controller.UserController
}

func main() {
	configVar, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db := config.NewDatabase(&configVar)

	fmt.Println("Connected to database")

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	appController := AppController{
		userController: userController,
	}

	router := NewRouter(appController)

	address := fmt.Sprintf("%v:%v", configVar.Host, configVar.Port)
	fmt.Println("Server started on " + address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
}
