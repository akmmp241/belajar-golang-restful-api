package service

import (
	"akmmp241/belajar-golang-restful-api/exception"
	"akmmp241/belajar-golang-restful-api/helper"
	"akmmp241/belajar-golang-restful-api/model/domain"
	"akmmp241/belajar-golang-restful-api/model/web"
	"akmmp241/belajar-golang-restful-api/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{categoryRepository: categoryRepository, DB: DB, Validate: validate}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.categoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.categoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.categoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.categoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.categoryRepository.Delete(ctx, tx, category)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.categoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	categories := service.categoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryListResponse(categories)
}
