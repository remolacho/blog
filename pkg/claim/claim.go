package claim

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
	jwt.StandardClaims
	ID int `json:"id"`
}

func (c *Claim) Encode(signingString string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(signingString))
}

func Decode(tokenString, signingString string) (*Claim, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(signingString), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claim")
	}

	iID, ok := claim["id"]
	if !ok {
		return nil, errors.New("user id not found")
	}

	id, ok := iID.(float64)
	if !ok {
		return nil, errors.New("invalid user id")
	}

	return &Claim{ID: int(id)}, nil
}
