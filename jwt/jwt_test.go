package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

var claims = Claims{
	Type:     "admin",
	SignKey:  []byte("secret"),
	Username: "zhl",
	RegisteredClaims: jwt.RegisteredClaims{
		Issuer:    "jwt",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	},
}
var token string

func TestClaims_Create(t *testing.T) {
	var err error
	token, err = claims.Create()
	if err != nil {
		t.Fatalf("claims.Create(): %v", err)
	}

	t.Log(token)
}

func TestClaims_Validate(t *testing.T) {
	t.Log(token)
	res, err := Validate(token, []byte("secret"))
	if err != nil {
		t.Fatalf("claims.Validate(): %v", err)
	}
	t.Log(res)
}
func TestClaims_Validate2(t *testing.T) {
	token2 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJUeXBlIjoiYWRtaW4iLCJTaWduS2V5IjoiYzJWamNtVjAiLCJVc2VybmFtZSI6InpobCIsImlzcyI6Imp3dCIsImV4cCI6MTczNDM0MDI2NCwibmJmIjoxNzM0MzM5OTY0LCJpYXQiOjE3MzQzMzk5NjR9.MxSEg2sTKPIrQjSIKy58uuJ0wVeUudUtbvHFQfHUPKE"
	res, err := Validate(token2, []byte("secret"))
	if err != nil {
		t.Fatalf("claims.Validate(): %v", err)
	}
	t.Log(res)

}
