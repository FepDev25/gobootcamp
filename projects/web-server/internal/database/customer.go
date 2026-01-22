package database

import (
	"context"
	"fmt"
	"web-server/internal/models"
)

func (s *service) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	query := `
		INSERT INTO customers (full_name, email, phone, address)
		VALUES ($1, $2, $3, $4)
		RETURNING customer_id, created_at
	`
	err := s.db.QueryRow(ctx, query,
		customer.FullName,
		customer.Email,
		customer.Phone,
		customer.Address,
	).Scan(&customer.ID, &customer.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}
	return nil
}

func (s *service) GetCustomer(ctx context.Context, id int) (*models.Customer, error) {
	query := `
		SELECT customer_id, full_name, email, phone, address, created_at
		FROM customers
		WHERE customer_id = $1
	`
	var c models.Customer
	err := s.db.QueryRow(ctx, query, id).Scan(
		&c.ID,
		&c.FullName,
		&c.Email,
		&c.Phone,
		&c.Address,
		&c.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}
	return &c, nil
}

func (s *service) GetAllCustomers(ctx context.Context) ([]models.Customer, error) {
	query := `
		SELECT customer_id, full_name, email, phone, address, created_at
		FROM customers
	`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get customers: %w", err)
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		if err := rows.Scan(
			&c.ID,
			&c.FullName,
			&c.Email,
			&c.Phone,
			&c.Address,
			&c.CreatedAt,
		); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (s *service) UpdateCustomer(ctx context.Context, customer *models.Customer) error {
	query := `
		UPDATE customers
		SET full_name = $1, email = $2, phone = $3, address = $4
		WHERE customer_id = $5
	`
	_, err := s.db.Exec(ctx, query,
		customer.FullName,
		customer.Email,
		customer.Phone,
		customer.Address,
		customer.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update customer: %w", err)
	}
	return nil
}

func (s *service) DeleteCustomer(ctx context.Context, id int) error {
	query := `DELETE FROM customers WHERE customer_id = $1`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}
	return nil
}
