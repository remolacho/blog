package users

import "blog/internal/entities/user"

type UserHandler struct {
	Repository user.Repository
}
