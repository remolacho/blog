package posts

import (
	"blog/pkg/response"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (handler *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = handler.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, response.Map{})
}
