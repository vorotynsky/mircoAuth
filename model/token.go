// Copyright (c) 2020 Vorotynsky Maxim

package model

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token string

type (
	jwtClaims struct {
		UserResource
		jwt.StandardClaims
	}

	JwtSecret jwt.Keyfunc
)

func NewJwtSecret(secret string) JwtSecret {
	return func(_ *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}
}

func (user User) GenerateToken(keyToken JwtSecret, lifeTime time.Duration) (Token, error) {
	user.DeletePrivates()
	jwtClaims := jwtClaims{
		UserResource: MakeResourse(user),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(lifeTime).Unix(),
		},
	}
	return makeToken(jwtClaims, keyToken)
}

func (tokenStr Token) DecryptToken(keyToken JwtSecret) (*User, error) {
	claims := &jwtClaims{}
	token, err := jwt.ParseWithClaims(string(tokenStr), claims, jwt.Keyfunc(keyToken))
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Token isn't valid. ")
	}
	user := RestoreResource(claims.UserResource)
	return &user, nil
}

func (tokenStr *Token) Refresh(keyToken JwtSecret, lifeTime time.Duration) error {
	claims := &jwtClaims{}
	if err := parseWithClaims(*tokenStr, claims, keyToken); err != nil {
		return err
	}

	claims.ExpiresAt = time.Now().Add(lifeTime).Unix()
	tokenString, err := makeToken(*claims, keyToken)
	if err != nil {
		return err
	}

	*tokenStr = Token(tokenString)
	return nil
}

func makeToken(claims jwtClaims, keyToken JwtSecret) (Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, _ := keyToken(nil)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return Token(tokenString), nil
}

func parseWithClaims(tokenStr Token, claims *jwtClaims, keyToken JwtSecret) error {
	token, err := jwt.ParseWithClaims(string(tokenStr), claims, jwt.Keyfunc(keyToken))
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("Token isn't valid. ")
	}
	return nil
}
