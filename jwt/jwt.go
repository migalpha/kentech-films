package jwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	film "github.com/migalpha/kentech-films"
	"github.com/migalpha/kentech-films/config"
)

var jwtSecret = []byte(config.Commons().JWTSecret)

// Claims is a struct to map a jwt token
type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
	jwt.StandardClaims
}

// GenerateToken Generate a new token
func GenerateToken(user film.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		user.ID.Uint(),
		user.Username.String(),
		user.IsActive,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Commons().AppName,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken return decoded JWT
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetToken return a string with JWT
func GetTokenSignature(c []string) (token string) {
	splitHeader := strings.Split(c[0], "Bearer")
	fullToken := strings.TrimSpace(splitHeader[1])
	splitToken := strings.Split(fullToken, ".")
	return splitToken[2]
}
