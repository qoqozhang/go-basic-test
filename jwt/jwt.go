package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Type     string
	SignKey  []byte
	Username string
	jwt.RegisteredClaims
}

func (c Claims) Create() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString(c.SignKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func Validate(tokenString string, SignKey []byte) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})
	switch {
	case token.Valid:
		fmt.Println("Validate successful")
		return true, nil
	case errors.Is(err, jwt.ErrTokenMalformed): // token格式解析错误
		fmt.Println("That's not even a token")
		return false, err
	case errors.Is(err, jwt.ErrTokenSignatureInvalid): // token的指纹校验失败
		fmt.Println("Invalid signature")
		return false, err
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet): // token已经过期
		fmt.Println("Timing is everything")
		return false, err
	default:
		fmt.Println("Couldn't handle this token")
		return false, err
	}
}
