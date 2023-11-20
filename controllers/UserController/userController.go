package UserController

import (
	"swisscast-go/services/AuthService"
	"swisscast-go/services/UserService"

	"github.com/gin-gonic/gin"
)


func Login(c* gin.Context) {
	AuthService.Login(c);
}

func Register(c* gin.Context) {
	AuthService.Register(c)
}

func SubscribeToFeed(c* gin.Context) {
	UserService.SubscribeToFeed(c)
}

func GetUserInfo(c* gin.Context) {
	AuthService.GetUserInfo(c)
}

func GetUserPodcasts(c *gin.Context) {
	UserService.GetUserPodcasts(c);
}