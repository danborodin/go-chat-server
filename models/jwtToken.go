package models

import "github.com/dgrijalva/jwt-go"

// Token ...
type Token struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	jwt.StandardClaims
}
