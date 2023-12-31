package models

import "time"

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
    Phone string `json:"phone"`
	DateOfBirth time.Time `json:"dateOfBirth"`
    Email string `json:"email"`
    Password string `json:"password"`
}

type UserPodcast struct {
	ID string `json:"id"`
	UserId string `json:"userId"`
	URL string `json:"url"`
	Username string `json:"username"`
	Pass string `json:"pass"`
	PodcastLogo string `json:"podcast_logo"`
	PodcastName string `json:"podcast_name"`
}