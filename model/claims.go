// Copyright (c) 2020 Vorotynsky Maxim

package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type ClaimType string

const (
	Role        = "Role"
	Password    = "Password"
	Permission  = "Permission"
	Restriction = "Restriction"
	UserName    = "UserName"
	Email       = "Email"
)

type Claim struct {
	Type  ClaimType
	Value ClaimValue
}

func NewStirngClaim(ctype ClaimType, value string) Claim {
	return Claim{ctype, StringValue{value}}
}

func NewPasswordClaim(password string) (Claim, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println("[PasswordHashing]:", err)
	}
	return Claim{Password, HashValue{hash}}, err
}

func (source Claim) HashPasswordClaim(password string) (Claim, error) {
	if _, ok := source.Value.(HashValue); ok {
		return source, nil
	}
	return NewPasswordClaim(source.Value.GetValue())
}
