package v1

import (
	"blog/internal/controllers/v1/posts"
	"blog/internal/controllers/v1/users"
	"blog/internal/entities/post"
	"blog/internal/entities/user"
	"blog/pkg/engineDB"
	"github.com/go-chi/chi"
	"net/http"
)

func New() http.Handler {
	router := chi.NewRouter()
	ur := &users.UserHandler{Repository: &user.UserService{Data: engineDB.Factory()}}
	pr := &posts.PostHandler{Repository: &post.PostService{Data: engineDB.Factory()}}

	router.Mount("/users", ur.GetRoutes())
	router.Mount("/posts", pr.GetRoutes())
	return router
}
