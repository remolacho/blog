package posts

import (
	"blog/internal/server/entities/post"
	"blog/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
)

func (handler *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p post.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = handler.Repository.Create(ctx, &p)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), p.ID))
	response.JSON(w, http.StatusCreated, response.Map{"post": p})
}
