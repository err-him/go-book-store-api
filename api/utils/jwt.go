package utils

import (
	c "book-store-api/api/constants"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// TokenClaims - Claims for a JWT access token.
type TokenClaims struct {
	Id    int64  `json:"userId"`
	Name  string `json:"userName"`
	Token string `json:"userKey"`
	jwt.StandardClaims
}

func IssueJWTToken(id int64, name, token string) (string, error) {
	//get the Jwt Security Key
	key, err := GetEnvVar(c.JWT_SECURITY_TOKEN)
	if err != nil {
		return "", err
	}
	var jwtKey = []byte(key)
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(c.JWT_EXPIRY_MINUTES * time.Minute).Unix()
	// Create the JWT claims, which includes the username and expiry time
	claims := &TokenClaims{
		Name:  name,
		Id:    id,
		Token: token,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString([]byte(jwtKey))
	return ("Bearer " + tokenString), err
}

func VerifyJWTToken(tknStr string) (*TokenClaims, error) {
	claims := &TokenClaims{}
	jwtKey, err := GetEnvVar(c.JWT_SECURITY_TOKEN)
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte(jwtKey), nil
	})
	if !token.Valid {
		return nil, err
	}
	return claims, err
}
