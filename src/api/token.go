package api

import (
	"time"

	"github.com/spf13/viper"
	"github.com/dgrijalva/jwt-go"

	"github.com/nirajgeorgian/account/src/model"
)

var (
	key = []byte(viper.GetString("jwtsecret"))
)

type CustomClaims struct {
	Account *model.Account
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(account *model.Account) (string, error)
}

// Encode a claim into a JWT
func (srv *AccountServer) Encode(account *model.Account) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		account,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "account.service",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}

// Decode a token string into a token object
func (srv *AccountServer) Decode(tokenString string) (*CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
