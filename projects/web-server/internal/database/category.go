package database

import (
	"context"
	"fmt"
	"web-server/internal/models"
)

func (s *service) CreateCategory(ctx context.Context, category *models.Category) error {
	query := `INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING category_id`
	err := s.db.QueryRow(ctx, query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}
	return nil
}

func (s *service) GetCategory(ctx context.Context, id int) (*models.Category, error) {
	query := `SELECT category_id, name, description FROM categories WHERE category_id = $1`
	var category models.Category
	err := s.db.QueryRow(ctx, query, id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}
	return &category, nil
}

func (s *service) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	query := `SELECT category_id, name, description FROM categories`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (s *service) UpdateCategory(ctx context.Context, category *models.Category) error {
	query := `UPDATE categories SET name = $1, description = $2 WHERE category_id = $3`
	_, err := s.db.Exec(ctx, query, category.Name, category.Description, category.ID)
	if err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}
	return nil
}

func (s *service) DeleteCategory(ctx context.Context, id int) error {
	query := `DELETE FROM categories WHERE category_id = $1`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil
}
