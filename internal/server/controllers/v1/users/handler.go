package users

import "blog/internal/server/entities/user"

type UserHandler struct {
	Repository user.Repository
}
