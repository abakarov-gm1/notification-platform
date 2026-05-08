package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Jwt struct {
	secret []byte
}

func NewJwt(secret string) *Jwt {
	return &Jwt{
		secret: []byte(secret),
	}
}

func (j *Jwt) GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // срок жизни
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.secret)
}

func (j *Jwt) DecodeToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return j.secret, nil
	})

	if err != nil {
		return 0, err
	}

	// достаём claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, jwt.ErrTokenInvalidClaims
	}

	// вытаскиваем user_id
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}

	return uint(userIDFloat), nil
}
