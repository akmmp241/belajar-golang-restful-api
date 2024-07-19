package app

import (
	"akmmp241/belajar-golang-restful-api/controller"
	"akmmp241/belajar-golang-restful-api/exception"
	"akmmp241/belajar-golang-restful-api/helper"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	router.GET("/ping", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		_, err := fmt.Fprint(writer, "pong")
		helper.PanicIfErr(err)
	})
	
	return router
}
