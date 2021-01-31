package response

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type Map map[string]interface{}

func JSON(response http.ResponseWriter, statusCode int, data interface{}) error {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.WriteHeader(statusCode)

	if data == nil {
		return nil
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response.Write(body)
	return nil
}

func HTTPError(response http.ResponseWriter, statusCode int, message string) error {
	msg := ErrorMessage{Message: message}

	return JSON(response, statusCode, msg)
}
