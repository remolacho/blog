package posts

import (
	"blog/internal/server/entities/post"
	"blog/pkg/response"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (handler *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	var p post.Post
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = handler.Repository.Update(ctx, uint(id), p)
	if err != nil {
		response.HTTPError(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
