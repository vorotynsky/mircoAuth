// Copyright (c) 2020 Vorotynsky Maxim

package model

import (
	"fmt"
)

type ClaimSet struct {
	Set map[ClaimType]ClaimValue
}

func NewClaimSet() ClaimSet {
	return ClaimSet{Set: map[ClaimType]ClaimValue{}}
}

func (claims ClaimSet) GetClaim(t ClaimType) (claim Claim, err error) {
	value, ok := claims.Set[t]
	if !ok {
		err = fmt.Errorf("Claim with type %v not found. ", t)
		value = StringValue{Value: "not found"}
	}
	claim = Claim{t, value}
	return
}

func (claims ClaimSet) SetClaim(claim Claim) {
	claims.Set[claim.Type] = claim.Value
}

func (claims ClaimSet) RemoveClaim(claim ClaimType) {
	delete(claims.Set, claim)
}

func (claims ClaimSet) GetClaims() []Claim {
	claimsArray := make([]Claim, 0, len(claims.Set))
	for key, value := range claims.Set {
		claimsArray = append(claimsArray, Claim{key, value})
	}
	return claimsArray
}
