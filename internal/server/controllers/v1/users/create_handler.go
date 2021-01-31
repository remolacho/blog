package users

import (
	"blog/internal/server/entities/user"
	"blog/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
)

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = handler.Repository.Create(ctx, &u)
	if err != nil {
		response.HTTPError(w, http.StatusBadRequest, err.Error())
		return
	}

	u.Password = ""

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), u.ID))
	response.JSON(w, http.StatusCreated, response.Map{"user": u})
}
