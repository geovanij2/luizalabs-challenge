package cryptography

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAdapter struct {
	secret []byte
}

func NewJwtAdapter(secret string) *JwtAdapter {
	return &JwtAdapter{
		secret: []byte(secret),
	}
}

type JwtClaims struct {
	ClientId string `json:"clientId"`
}

func (j *JwtAdapter) Encrypt(plainText string) (string, error) {
	claims := jwt.MapClaims{
		"clientId": plainText,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JwtAdapter) Decrypt(encryptedText string) (string, error) {
	token, err := jwt.Parse(encryptedText, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return j.secret, nil
	})

	if err != nil {
		fmt.Println("Erro ao validar o token:", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["clientId"].(string), nil
	} else {
		return "", fmt.Errorf("token inválido")
	}
}
