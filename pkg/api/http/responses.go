package http

import "time"

type User struct {
	Id        string     `json:"id"`
	FirstName *string    `json:"first_name,omitempty"`
	LastName  *string    `json:"last_name,omitempty"`
	Nickname  *string    `json:"nickname,omitempty"`
	Email     string     `json:"email"`
	Country   *string    `json:"country,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
