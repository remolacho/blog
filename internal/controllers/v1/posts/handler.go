package posts

import "blog/internal/entities/post"

type PostHandler struct {
	Repository post.Repository
}
