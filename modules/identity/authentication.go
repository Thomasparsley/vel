package identity

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func Login(username string, password string) {

}

func Verify(verifedToken []byte, tokenToVerify string) {
	tokenToVerify = strings.TrimSpace(tokenToVerify)

	if tokenToVerify == "" {
		return // TODO:
	}

	token, err := jwt.Parse(tokenToVerify, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(
				"unexpected signing method: %v", t.Header["alg"])
		}

		return verifedToken, nil
	})
	if err != nil {
		return // TODO:
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return // TODO:
	}

	userIdString := claims["iss"].(string)
	expiresAt := int64(claims["exp"].(float64))

	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		return // TODO:
	}

	if userId == 0 {
		return // TODO:
	}
	if expiresAt == 0 {
		return // TODO:
	}

	var newToken string
	if time.Unix(expiresAt, 0).Sub(time.Now()).Minutes() < 45 {
		newClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    userIdString,
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		})

		newToken, err = newClaims.SignedString(verifedToken)
		if err != nil {
			return // TODO:
		}
	}

	if newToken == "" {
		return // TODO:
	}

	return // TODO:
}
