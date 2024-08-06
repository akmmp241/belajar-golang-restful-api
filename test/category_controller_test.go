package test

import (
	"akmmp241/belajar-golang-restful-api/app"
	"akmmp241/belajar-golang-restful-api/controller"
	"akmmp241/belajar-golang-restful-api/helper"
	"akmmp241/belajar-golang-restful-api/middleware"
	"akmmp241/belajar-golang-restful-api/model/domain"
	"akmmp241/belajar-golang-restful-api/repository"
	"akmmp241/belajar-golang-restful-api/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "joko:akuanakhebat@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)

	_, _ = db.Exec("TRUNCATE TABLE categories")

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setupRouter(setupTestDB())

	requestBody := strings.NewReader(`{"name":"test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusCreated), responseBody["status"])
	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	router := setupRouter(setupTestDB())

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusBadRequest), responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository().Save(context.Background(), tx, domain.Category{
		Name: "test",
	})
	_ = tx.Commit()

	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(categoryRepository.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]any)["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+"4", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusNotFound), responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test1",
	})
	_ = tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)
	log.Println(responseBody["data"])

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody["data"])

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusNotFound), responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test1",
	})
	_ = tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)
	log.Println(responseBody["data"])

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody["data"])

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusNotFound), responseBody["status"])
}

func TestGetListCategorySuccess(t *testing.T) {
	db := setupTestDB()

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test1",
	})
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test2",
	})
	_ = tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)
	log.Println(responseBody["data"])

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), responseBody["status"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test1",
	})
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test2",
	})
	_ = tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-API-KEY", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	_ = json.Unmarshal(body, &responseBody)
	log.Println(responseBody["data"])

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), responseBody["status"])
}
