package users

import (
	"blog/internal/entities/user"
	"blog/pkg/claim"
	"blog/pkg/response"
	"encoding/json"
	"net/http"
	"os"
)

func (handler *UserHandler) Login(responseWriter http.ResponseWriter, request *http.Request) {
	var user user.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.HTTPError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	storedUser, err := handler.Repository.FindByUsername(user.Username)
	if err != nil {
		response.HTTPError(responseWriter, http.StatusNotFound, err.Error())
		return
	}

	if !storedUser.PasswordMatch(user.Password) {
		response.HTTPError(responseWriter, http.StatusBadRequest, "password don't match")
		return
	}

	clain := claim.Claim{ID: int(storedUser.ID)}
	token, err := clain.Encode(os.Getenv("SIGNING_STRING"))
	if err != nil {
		response.HTTPError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	responseWriter.Header().Set("Authorization", "Bearer "+token)

	response.JSON(responseWriter, http.StatusOK, response.Map{"token": token})
}
