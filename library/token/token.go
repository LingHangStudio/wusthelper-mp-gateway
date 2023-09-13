package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Token struct {
	SecretKey string
	Timeout   time.Duration
}

func New(secret string, timeout time.Duration) *Token {
	return &Token{
		SecretKey: secret,
		Timeout:   timeout,
	}
}

func (t *Token) Sign(oid string) (token string) {
	//fmt.p
	claims := &jwt.MapClaims{
		"openid": oid,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Unix() + int64(t.Timeout.Seconds()),
	}
	token, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, *claims).
		SignedString([]byte(t.SecretKey))
	return
}

func (t *Token) Verify(token string) bool {
	tt, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})

	return err == nil && tt.Valid
}

func (t *Token) GetClaim(token string) bool {
	tt, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})

	return err == nil && tt.Valid
}

func (t *Token) GetClaimWithoutVerify(token string) *jwt.MapClaims {
	claims := new(jwt.MapClaims)
	_, _, err := jwt.NewParser().ParseUnverified(token, claims)
	if err != nil {
		fmt.Println("GetClaimWithoutVerifyError")
		fmt.Println(err.Error())
		return nil
	}

	return claims
}
