package UserService

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"swisscast-go/controllers/FeedController"
	"swisscast-go/db"
	"swisscast-go/models"
	"swisscast-go/utils/AppUtils"
	"swisscast-go/utils/JwtUtils"

	"github.com/gin-gonic/gin"
)

type SubscribeToFeedParams struct {
	Username string `json:"username"`
	Pass  string `json:"pass"`
	Url   string `json:"url"`
}

type GetUserPodcastsParams struct {
	UserId string `json:"userId"`
}

func SubscribeToFeed(c *gin.Context) bool{
	conn, err := db.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()


	var subscribeParams FeedController.RSSFeedRequestParams
	c.BindJSON(&subscribeParams)

	reqToken := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
	claims, err := JwtUtils.DecryptToken(reqToken)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Seu token expirou, faça login novamente"});
		return false
	}

	hashedPassword, err := AppUtils.HashPassword(subscribeParams.Pass)
	if err != nil {
		log.Fatal(err)
		return false
	}

	feed := FeedController.FetchRssFeedData(c, subscribeParams)

	var user_id string
	var url string
	existsErr := conn.QueryRow("SELECT user_id, url from user_podcasts WHERE user_id = $1 and url = $2", claims.User.ID, subscribeParams.Url).Scan(&user_id, &url)
	if existsErr != nil {
        if existsErr != sql.ErrNoRows  {
            log.Print(existsErr)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Ocorreu um erro, tente novamente"})
			return false
        }

		fmt.Println("does not exist")

		_, insertErr := conn.Exec("INSERT INTO user_podcasts(user_id, url, username, pass, podcast_logo, podcast_name) VALUES ($1, $2, $3, $4, $5, $6)", claims.User.ID, subscribeParams.Url, subscribeParams.Username, hashedPassword, feed.Image.URL, feed.Title)
		if insertErr != nil {
			log.Print(insertErr)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro ao inserir, tente novamente"})
			return false
		}

		c.IndentedJSON(http.StatusOK, gin.H{"feed": feed})
		return true
    }

	c.IndentedJSON(http.StatusOK, gin.H{"feed": feed})
	return true
}

func GetUserPodcasts(c *gin.Context){
	conn, err := db.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reqToken := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
	claims, err := JwtUtils.DecryptToken(reqToken)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Seu token expirou, faça login novamente"});
		return
	}

	podcasts, queryErr := conn.Query("SELECT id, podcast_logo, podcast_name, url from user_podcasts WHERE user_id = $1", claims.User.ID)
	if queryErr != nil {
		if queryErr != sql.ErrNoRows  {
			log.Print(queryErr)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Ocorreu um erro"})
			return
        }
	}
	
	var user_podcasts []models.UserPodcast
	for podcasts.Next() {
       var user_podcast models.UserPodcast
	   err = podcasts.Scan(&user_podcast.ID, &user_podcast.PodcastLogo, &user_podcast.PodcastName, &user_podcast.URL)
	   if err != nil {
		continue
	   }

	   user_podcasts = append(user_podcasts, user_podcast)
	}
	
    c.IndentedJSON(http.StatusOK, gin.H{"podcasts": user_podcasts}) 
}