package JwtUtils

import (
	"fmt"
	"swisscast-go/db"
	"swisscast-go/models"
	"swisscast-go/utils/AppUtils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey string = AppUtils.GoDotEnvVariable("SECRET_KEY")

func GenerateToken(username string) (string, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	var userId string
	userErr := conn.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userId)
	if userErr != nil {
		return "", userErr
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.Claims{
		Id: userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecryptToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
