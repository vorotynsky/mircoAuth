// Copyright (c) 2020 Vorotynsky Maxim

package controllers

import "microAuth/model"

type userResource struct {
	UserID model.UserID               `json:"user_id"`
	Claims map[model.ClaimType]string `json:"claims"`
}

func makeResourse(usr model.User) userResource {
	usr.Claims.RemoveClaim(model.Password)
	claims := make(map[model.ClaimType]string)
	for k, v := range usr.Claims.Set {
		claims[k] = v.GetValue()
	}
	return userResource{usr.UserID, claims}
}

func restoreResource(res userResource) model.User {
	claims := model.NewClaimSet()
	for k, v := range res.Claims {
		claim := model.NewStirngClaim(k, v)
		claims.SetClaim(claim)
	}
	return model.User{res.UserID, claims}
}
