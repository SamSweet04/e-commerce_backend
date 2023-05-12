package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

func GenerateToken(userId int) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Hour * 1).Unix()
	claims := &JWTClaim{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
func ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return "", err
	}
	return strconv.FormatInt(int64(claims.UserID), 10), nil
}
