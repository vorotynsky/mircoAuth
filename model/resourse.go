// Copyright (c) 2020 Vorotynsky Maxim

package model

type UserResource struct {
	UserID UserID               `json:"user_id"`
	Claims map[ClaimType]string `json:"claims"`
}

func MakeResourse(usr User) UserResource {
	usr.DeletePrivates()
	claims := make(map[ClaimType]string)
	for k, v := range usr.Claims.Set {
		claims[k] = v.GetValue()
	}
	return UserResource{usr.UserID, claims}
}

func RestoreResource(res UserResource) User {
	claims := NewClaimSet()
	for k, v := range res.Claims {
		claim := NewStirngClaim(k, v)
		claims.SetClaim(claim)
	}
	return User{res.UserID, claims}
}
