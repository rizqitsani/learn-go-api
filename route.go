package main

import "github.com/julienschmidt/httprouter"

func NewRouter(appController AppController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", appController.userController.FindAll)
	router.GET("/api/users/:id", appController.userController.FindById)
	router.POST("/api/users", appController.userController.Create)

	return router
}
