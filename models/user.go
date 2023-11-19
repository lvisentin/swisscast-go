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

