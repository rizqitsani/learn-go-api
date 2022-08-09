package main

import "github.com/julienschmidt/httprouter"

func NewRouter(appController AppController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", appController.userController.FindAll)
	router.GET("/api/users/:id", appController.userController.FindById)
	router.POST("/api/users", appController.userController.Create)
	router.PUT("/api/users/:id", appController.userController.Update)
	router.DELETE("/api/users/:id", appController.userController.Delete)

	return router
}
