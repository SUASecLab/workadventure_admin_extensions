package extensions

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func DecodeToken(jwtToken, externalToken string) (bool, map[string]interface{}) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: failed to decode token: unexpected signing method: %v",
				time.Now().Format(time.Stamp), token.Header["alg"])
		}

		return []byte(externalToken), nil
	})

	if err != nil {
		log.Printf("%s: failed to parse token: %s", time.Now().Format(time.Stamp), err)
		return false, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims
	} else {
		log.Printf("%s Invalid token: %s", time.Now().Format(time.Stamp), err)
	}

	return false, nil
}
