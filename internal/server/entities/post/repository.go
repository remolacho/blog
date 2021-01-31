package post

import "context"

// Repository handle the CRUD operations with Posts.
type Repository interface {
	All(ctx context.Context) ([]Post, error)
	Find(ctx context.Context, id uint) (Post, error)
	FindByUser(ctx context.Context, userID uint) ([]Post, error)
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, id uint, post Post) error
	Delete(ctx context.Context, id uint) error
	logger(query string, params []string)
}
