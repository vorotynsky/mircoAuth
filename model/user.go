// Copyright (c) 2020 Vorotynsky Maxim

package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserID int32

type User struct {
	UserID UserID
	Claims ClaimSet
}

func NewUser(id UserID, name string, password string) *User {
	claimSet := NewClaimSet()
	claimSet.SetClaim(NewStirngClaim(UserName, name))

	passC, err := NewPasswordClaim(password)
	if err != nil {
		return nil
	}
	claimSet.SetClaim(passC)

	return &User{id, claimSet}
}

func NewUserWithHashedPass(id UserID, name string, password []byte) User {
	claimSet := NewClaimSet()
	claimSet.SetClaim(NewStirngClaim(UserName, name))
	claimSet.SetClaim(Claim{Password, HashValue{password}})

	return User{id, claimSet}
}

func (usr *User) ConfirmPassword(password string) error {
	if usr == nil {
		return errors.New("user is null ")
	}
	claim, err := usr.Claims.GetClaim(Password)
	if err != nil {
		return err
	}
	if hash, ok := claim.Value.(HashValue); !ok {
		return bcrypt.CompareHashAndPassword(hash.Hash, []byte(password))
	}
	return errors.New("Password claim doesn't contains password hash. ")
}

func (usr *User) HashUserPassword() ([]byte, error) {
	passClaim, err := usr.Claims.GetClaim(Password)
	if err != nil {
		return nil, err
	}
	hashVal, ok := passClaim.Value.(HashValue)
	if !ok {
		newPassCalim, err := NewPasswordClaim(passClaim.Value.GetValue())
		if err != nil {
			return nil, err
		}
		usr.Claims.SetClaim(newPassCalim)
		passClaim, err = usr.Claims.GetClaim(Password)
		if err != nil {
			return nil, err
		}
		if hashVal, ok = passClaim.Value.(HashValue); !ok {
			err = errors.New("[HashUserPassword]: post hash error. ")
			return nil, err
		}
	}
	return hashVal.Hash, nil
}
