package posts

import "blog/internal/server/entities/post"

type PostHandler struct {
	Repository post.Repository
}
