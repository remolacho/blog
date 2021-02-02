package users

import (
	"blog/pkg/response"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (handler *UserHandler) All(w http.ResponseWriter, r *http.Request) {
	users, err := handler.Repository.All()
	if err != nil {
		response.HTTPError(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, response.Map{"users": users})
}

func (handler *UserHandler) Find(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	u, err := handler.Repository.Find(uint(id))
	if err != nil {
		response.HTTPError(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, response.Map{"user": u})
}
