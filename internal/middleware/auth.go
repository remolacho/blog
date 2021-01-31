package middleware

import (
	"blog/pkg/claim"
	"blog/pkg/response"
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
)

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("autorization is required")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("invalid autorization format")
	}

	token := strings.Split(authorization, " ")
	if len(token) != 2 {
		return "", errors.New("invalid autorization format")
	}

	return token[1], nil
}

func Authorizator(next http.Handler) http.Handler {
	signingString := os.Getenv("SIGNING_STRING")
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		authorization := request.Header.Get("Authorization")
		tokenString, err := tokenFromAuthorization(authorization)
		if err != nil {
			response.HTTPError(responseWriter, http.StatusUnauthorized, err.Error())
			return
		}

		decode, err := claim.Decode(tokenString, signingString)
		if err != nil {
			response.HTTPError(responseWriter, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := request.Context()
		ctx = context.WithValue(ctx, "id", decode.ID)
		next.ServeHTTP(responseWriter, request.WithContext(ctx))
	})
}
