package FeedController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

type RSSFeedRequestParams struct {
	Username string `json:"username"`
	Pass  string `json:"pass"`
	Url   string `json:"url"`
}

func GetRssFeed(c *gin.Context) {
	var rssFeedParams RSSFeedRequestParams
	c.BindJSON(&rssFeedParams)

	fp := gofeed.NewParser()
	fp.AuthConfig = &gofeed.Auth{
		Username: rssFeedParams.Username,
		Password: rssFeedParams.Pass,
	  }
	feed, _ := fp.ParseURL(rssFeedParams.Url)
	c.JSON(http.StatusOK, gin.H{"feed":feed})
}
