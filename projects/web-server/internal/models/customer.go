package models

import "time"

type Customer struct {
	ID        int       `json:"customer_id"`
	FullName  string    `json:"full_name"`
	Email     *string   `json:"email"`
	Phone     *string   `json:"phone"`
	Address   *string   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
