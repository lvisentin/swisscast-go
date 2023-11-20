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

	var user models.User
	userErr := conn.QueryRow("SELECT id, username, name, phone, dateofbirth, email, password FROM users WHERE username = $1", username).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Phone,
		&user.DateOfBirth,
		&user.Email,
		&user.Password,
	)
	
	if userErr != nil {
		fmt.Printf("user not found", userErr)
		return "", userErr
	}

	fmt.Printf(user.Username, user.ID, user.Email)

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.Claims{
		User: user,
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

	fmt.Printf(token.Raw)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
