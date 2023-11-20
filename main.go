package main

import (
	"fmt"
	"swisscast-go/controllers/UserController"
	"swisscast-go/utils/AppUtils"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}


func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

    router.GET("/user", UserController.GetUserInfo)
    router.GET("/userPodcasts", UserController.GetUserPodcasts)
    
    router.POST("/login", UserController.Login)
    router.POST("/register", UserController.Register)
    router.POST("/subscribe", UserController.SubscribeToFeed)
	
    // router.POST("/feed", FeedController.GetRssFeed)
	router.Run(fmt.Sprintf("%s:%s", AppUtils.GoDotEnvVariable("GO_URL"), AppUtils.GoDotEnvVariable("GO_PORT")))
}
