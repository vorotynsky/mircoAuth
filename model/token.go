// Copyright (c) 2020 Vorotynsky Maxim

package model

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token string

type jwtClaims struct {
	UserResource
	jwt.StandardClaims
}

func (user User) GenerateToken() (Token, error) {
	user.DeletePrivates()
	jwtClaims := jwtClaims{
		UserResource: MakeResourse(user),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	secret, _ := keyToken(nil)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return Token(tokenString), nil
}

func (tokenStr Token) DecryptToken() (*User, error) {
	claims := &jwtClaims{}
	token, err := jwt.ParseWithClaims(string(tokenStr), claims, keyToken)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Token isn't valid. ")
	}
	user := RestoreResource(claims.UserResource)
	return &user, nil
}

//for fix import cycle
var keyToken func(_ *jwt.Token) (interface{}, error)

func SetSecretToken(secret string) {
	keyToken = func(_ *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}
}
