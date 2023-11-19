package UserService

import (
	"log"
	"strings"

	"swisscast-go/controllers/FeedController"
	"swisscast-go/db"
	"swisscast-go/utils/JwtUtils"

	"github.com/gin-gonic/gin"
)

type SubscribeToFeedParams struct {
	Username string `json:"username"`
	Pass  string `json:"pass"`
	Url   string `json:"url"`
}

func SubscribeToFeed(c *gin.Context) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()


	var subscribeParams FeedController.RSSFeedRequestParams
	c.BindJSON(&subscribeParams)

	reqToken := strings.Split(c.Request.Header["Authorization"][0], " ")[0]
	claims, err := JwtUtils.DecryptToken(reqToken)
	if err != nil {
		log.Fatal(err)
	}

	conn.QueryRow("INSERT INTO user_podcasts(user_id, url) VALUES ($1, $2)", claims.Id, subscribeParams.Url)

	FeedController.GetRssFeed(c, subscribeParams)
}