package utils

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// 盐
	jwtSecret = []byte("gV5Af9fJohYC")
	// 发行人
	issur = "Tamanya"

	// ErrorExpired 过期
	ErrorExpired = errors.New("token is expired")
	// ErrorNotValidYet 未生效
	ErrorNotValidYet = errors.New("token not active yet")
	// ErrorMalformed 畸形Token
	ErrorMalformed = errors.New("that's not even a token")
	// ErrorInvalid 其他错误
	ErrorInvalid = errors.New("couldn't handle this token")
)

// CustomClaims 定义一个Claims结构体,可以添加自定义属性
type CustomClaims struct {
	UserID       int64 `json:"user_id"`
	TokenVersion int64 `json:"token_version"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(userID, tokenVersion int64) (string, error) {
	// 新建一个Claims
	claims := CustomClaims{
		userID,
		tokenVersion,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(Config.TokenExpires) * time.Second).Unix(), // 设置超时时间
			Issuer:    issur,                                                                   // 设置发行人
		},
	}
	// 生成Token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签署
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrorMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrorExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrorNotValidYet
			} else {
				return nil, ErrorInvalid
			}
		}
	}
	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, ErrorInvalid
}
