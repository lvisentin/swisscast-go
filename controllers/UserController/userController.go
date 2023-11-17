package UserController

import (
	"swisscast-go/services/UserService"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
    Phone string `json:"phone"`
	DateOfBirth time.Time `json:"dateOfBirth"`
    Email string `json:"email"`
    Password string `json:"password"`
}

func Login(c* gin.Context) {
	UserService.Login(c);
}