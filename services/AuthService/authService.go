package AuthService

import (
	"log"
	"net/http"
	"swisscast-go/db"
	"swisscast-go/models"
	"swisscast-go/utils/AppUtils"
	"swisscast-go/utils/JwtUtils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var secretKey string = AppUtils.GoDotEnvVariable("SECRET_KEY")


func Login(c *gin.Context) {
	var user models.LoginUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if isValid := validateUserCredentials(user.Username, user.Password); !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := JwtUtils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) (user models.User, insertErr error) {
	var registerUser models.RegisterUser
	if err := c.BindJSON(&registerUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	hashedPassword, err := AppUtils.HashPassword(registerUser.Password)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	sql := `INSERT INTO users (name, phone, dateofbirth, email, password, username, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7);`
	
	insertErr = conn.QueryRow(sql, registerUser.Name, registerUser.Phone, registerUser.DateOfBirth, registerUser.Email, hashedPassword, registerUser.Username, time.Now()).Scan(&user)

	return user, insertErr
}

func validateUserCredentials(username, password string) bool {
	conn, err := db.OpenConnection()
	if err != nil {
		return false
	}
	defer conn.Close()

	var storedPassword string
	validateErr := conn.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
	if validateErr != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}