package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const jwtSecret string = "jdnfksdmfksd"

// CreateToken generates a new jwt token
func CreateToken(userid uint) (string, error) {
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
