package helper

import (
	"akmmp241/belajar-golang-restful-api/model/domain"
	"akmmp241/belajar-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryListResponse(categories []domain.Category) []web.CategoryResponse {
	var categoriesResponse []web.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}
	return categoriesResponse
}
