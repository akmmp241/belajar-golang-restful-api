package controller

import (
	"akmmp241/belajar-golang-restful-api/helper"
	"akmmp241/belajar-golang-restful-api/model/web"
	"akmmp241/belajar-golang-restful-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{CategoryService: categoryService}
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryCreateRequest web.CategoryCreateRequest
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.Response{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryUpdateRequest web.CategoryUpdateRequest
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfErr(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfErr(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfErr(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	categoriesResponse := controller.CategoryService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}
