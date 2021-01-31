package posts

import (
	"blog/pkg/response"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (handler *PostHandler) All(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := handler.Repository.All(ctx)
	if err != nil {
		response.HTTPError(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, response.Map{"posts": posts})
}

func (handler *PostHandler) Find(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	p, err := handler.Repository.Find(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, response.Map{"post": p})
}

func (handler *PostHandler) FindByUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userId")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	posts, err := handler.Repository.FindByUser(ctx, uint(userID))
	if err != nil {
		response.HTTPError(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, response.Map{"posts": posts})
}
