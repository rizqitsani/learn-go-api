package main

import "github.com/julienschmidt/httprouter"

func NewRouter(appController AppController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", appController.userController.Create)

	return router
}
