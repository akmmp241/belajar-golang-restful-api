package repository

import (
	"akmmp241/belajar-golang-restful-api/helper"
	"akmmp241/belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := `INSERT INTO categories(name) VALUES (?)`
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := `UPDATE categories SET name = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfErr(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := `DELETE FROM categories WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfErr(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := `SELECT id, name FROM categories WHERE id = ?`
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfErr(err)
	defer rows.Close()

	var category domain.Category
	if !rows.Next() {
		return category, errors.New("category not found")
	}

	err = rows.Scan(&category.Id, &category.Name)
	helper.PanicIfErr(err)
	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := `SELECT id, name FROM categories`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)
		categories = append(categories, category)
	}

	return categories
}
