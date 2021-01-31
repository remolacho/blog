package v1

import (
	"blog/internal/engineDB"
	"blog/internal/server/controllers/v1/posts"
	"blog/internal/server/controllers/v1/users"
	"blog/internal/server/entities/post"
	"blog/internal/server/entities/user"
	"github.com/go-chi/chi"
	"net/http"
)

func New() http.Handler {
	router := chi.NewRouter()
	ur := &users.UserHandler{Repository: &user.UserRepository{Data: engineDB.New()}}
	pr := &posts.PostHandler{Repository: &post.PostRepository{Data: engineDB.New()}}

	router.Mount("/users", ur.GetRoutes())
	router.Mount("/posts", pr.GetRoutes())
	return router
}
