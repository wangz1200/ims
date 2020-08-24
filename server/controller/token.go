package controller

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	tokenKey = []byte("AllYourBase")
)

func buildToken(user string) (string, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.StandardClaims{
			Issuer: user,
		})
	return t.SignedString(tokenKey)
}

func parseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(
		token,
		func(t *jwt.Token) (i interface{}, err error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Token令牌解析方法错误！")
			} else {
				return tokenKey, nil
			}
		})
}