package models

import "github.com/golang-jwt/jwt/v5"

//karna jwt gak bisa gini doang jadi butuh yang namanya registered claim
//registered claim merupakan claim2 standar / claim2 umum yang wajib di embed ditoken jwt yang ingin di publish
type AuthClaimJWT struct {
	jwt.RegisteredClaims
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}
