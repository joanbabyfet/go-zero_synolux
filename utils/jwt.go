package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	UserID   string `json:"user_id"`  //用户id
	Username string `json:"username"` //用户名
	//Role     int    `json:"role"` // 用户权限 1普通用户; 2管理员;
}

type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

// GetToken 生成token
func GetToken(user JwtPayLoad, accessSecret string, expires int64) (string, error) {
	claims := CustomClaims{
		JwtPayLoad: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expires))),
			Issuer:    "api_study",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

// 解析token
func ParseToken(tokenString string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
