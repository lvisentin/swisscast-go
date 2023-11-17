package UserService

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"swisscast-go/db"

	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`

}

const secretKey = "your-secret-key"

func Login(c *gin.Context) {
	var user LoginUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if isValid := validateUserCredentials(user.Username, user.Password); !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := generateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func validateUserCredentials(username, password string) bool {
	conn, err := db.OpenConnection()
	if err != nil {
		return false
	}

	var storedPassword string
	validateErr := conn.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
	fmt.Print("stooredPassword1 ", storedPassword)
	fmt.Print("stooredPassword1 ", validateErr)
	if validateErr != nil {
		return false
	}

	return storedPassword == password
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
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