package jwt

import "github.com/golang-jwt/jwt/v5"

func DecodeToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return "", nil
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
