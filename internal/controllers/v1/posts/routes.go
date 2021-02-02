package posts

import (
	"blog/internal/middleware"
	"github.com/go-chi/chi"
	"net/http"
)

func (handler *PostHandler) GetRoutes() http.Handler {
	router := chi.NewRouter()

	router.With(middleware.Authorizator).Get("/user/{userId}", handler.FindByUser)
	router.With(middleware.Authorizator).Get("/", handler.All)
	router.With(middleware.Authorizator).Post("/", handler.Create)
	router.With(middleware.Authorizator).Get("/{id}", handler.Find)
	router.With(middleware.Authorizator).Put("/{id}", handler.Update)
	router.With(middleware.Authorizator).Delete("/{id}", handler.Delete)
	return router
}
