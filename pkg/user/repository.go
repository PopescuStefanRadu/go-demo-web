package user

import (
	"context"
	"faceit/pkg/query"
)

type Repository interface {
	Migrate(ctx context.Context) error
	FindUsers(ctx context.Context, filter *FilterQuery) (*query.PaginatedResponse[User], error)
	DeleteUser(ctx context.Context, userID string) (*User, error)
	AddUser(ctx context.Context, addUserInput User) (User, error)
	UpdateUser(ctx context.Context, input *UpdateUserInput) (*User, error)
}
