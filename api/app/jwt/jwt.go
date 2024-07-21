package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	BearerTokenType        = "bearer"
	RefreshTokenType       = "refresh-token"
	PasswordResetTokenType = "password-reset"
)

type ToMapClaims interface {
	ToMapClaims() (jwt.MapClaims, error)
}

func anyToMapClaims(value any) (jwt.MapClaims, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	var result jwt.MapClaims
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type Claims struct {
	Type             string `json:"type" validate:"required"`
	Subject          string `json:"sub" validate:"required"`
	NotBeforeSeconds int64  `json:"nbf" validate:"required"`
	IssuedAtSeconds  int64  `json:"iat" validate:"required"`
	Issuer           string `json:"iss" validate:"required"`
	ExpiresAtSeconds int64  `json:"exp" validate:"required"`
}

func (claims *Claims) ToMapClaims() (jwt.MapClaims, error) {
	return anyToMapClaims(claims)
}

func (claims *Claims) ToRefreshClaims(expiresInSeconds int64) *Claims {
	claims.Type = RefreshTokenType
	claims.ExpiresAtSeconds = expiresInSeconds
	return claims
}

func ParseClaimsFromToken[C Claims](tokenString string, privateKey string) (*C, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		return nil, err
	}
	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		bytes, err := json.Marshal(&mapClaims)
		if err != nil {
			return nil, err
		}
		var claims C
		if err := json.Unmarshal(bytes, &claims); err != nil {
			return nil, err
		}
		return &claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func ParseClaimsFromTokenNoValidation[C Claims](tokenString string) (*C, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token")
	}
	bytes, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	var claims C
	if err := json.Unmarshal(bytes, &claims); err != nil {
		return nil, err
	}
	return &claims, nil
}

func CreateToken[C ToMapClaims](claims C, privateKey string) (string, error) {
	mapClaims, err := claims.ToMapClaims()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
