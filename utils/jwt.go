package utils

import (
	"errors"
	"ginchat/global"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte(global.MY_CONFIG.JWT.SiginKey)

type MyCustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	Name  string
	Email string
	UUID  string
}

// 创建 jwt
func CreateClaims(baseClaims BaseClaims) MyCustomClaims {
	claims := MyCustomClaims{
		baseClaims,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    global.MY_CONFIG.JWT.Issuer,
		},
	}

	return claims
}

func CreateJWT(claims MyCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

// 解析 jwt
var (
	TokenExpired     = errors.New("Token is expired") // 超时
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func ParseJWT(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
