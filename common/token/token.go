package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"time"
)

var (
	key []byte = []byte("jwt-jlang@jlang.tv.com")
)

//jwt
// 产生json web token
func GenToken(userID string) string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: time.Now().Add(24 * time.Hour * time.Duration(1)).Unix(),
		Issuer:    "jlang*yyh",
		Id:        userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		log.Error("create token err:", err)
		return ""
	}
	return ss
}

// 校验token是否有效
func CheckToken(token string) ( string, bool) {
	tokenstruct, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		log.Error("parase with claims failed.", err)
		return "", false
	}
	claims := tokenstruct.Claims.(jwt.MapClaims)
	id:= claims["jti"].(string)
	return id, true
}