package user

import "context"

// Repository handle the CRUD operations with Users.
type Repository interface {
	All(ctx context.Context) ([]User, error)
	Find(ctx context.Context, id uint) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id uint, user User) error
	Delete(ctx context.Context, id uint) error
	logger(query string, params []string)
}
