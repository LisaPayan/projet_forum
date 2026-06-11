package auth

import (
	"projet_forum/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	userid string 'json:userid'
	role string 'json:role'
	jwt.RegisteredClaims
}

func GenerateToken(userid string , role string) (string, error) {
	now:=time.Now()
	secret :=[]byte(config.getenvwithdefault("JWT_Secret","secret_forum)

	claims := claims{
		userid: userid,
		role: role,
		RegisteredClaims: jwt.RegisteredClaims{	
			subject: userid,
			issuer: "forum_api",
			audience: []string{"forum_front"},
			issuedat: jwt.NewNumericDate(now),
			expiresat: jwt.NewNumericDate(now.Add(15*time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}