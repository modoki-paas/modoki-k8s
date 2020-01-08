package jwtutil

import (
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/xerrors"
)

type JWT struct {
	key []byte

	SigningMethod jwt.SigningMethod
}

func NewJWT(key string) *JWT {
	return &JWT{
		key:           []byte(key),
		SigningMethod: jwt.SigningMethodRS256,
	}
}

func (j *JWT) issue(mapClaims jwt.MapClaims) (string, error) {
	tok := jwt.New(j.SigningMethod)

	tok.Claims = mapClaims

	token, err := tok.SignedString(j.key)

	if err != nil {
		return "", xerrors.Errorf("failed to sign key: %w", err)
	}

	return token, nil

}

func (j *JWT) validate(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if reflect.TypeOf(j.SigningMethod) != reflect.TypeOf(token.Method) {
			return nil, xerrors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return j.SigningMethod, nil
	})

	if err != nil {
		return nil, xerrors.Errorf("failed to parse token: %w", err)
	}

	return token, nil
}

func (j *JWT) IssueStateToken(state string) (string, error) {
	claims := jwt.MapClaims{}
	claims["state"] = state

	token, err := j.issue(claims)

	if err != nil {
		return "", xerrors.Errorf("failed to issue token: %w", err)
	}

	return token, nil
}

func (j *JWT) ValidateStateToken(tokenString, expectedState string) bool {
	token, err := j.validate(tokenString)

	if err != nil {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}

	if state, ok := claims["state"].(string); !ok {
		return false
	} else if state != expectedState {
		return false
	}

	return true
}
