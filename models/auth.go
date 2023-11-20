package models

import "github.com/dgrijalva/jwt-go"

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Username string `json:"username"`
	Name string `json:"name"`
    Phone string `json:"phone"`
	DateOfBirth string `json:"dateOfBirth"`
    Email string `json:"email"`
    Password string `json:"password"`
}

type Claims struct {
	Id string `json:"id"`
	Username string `json:"username"`
	User User `json:"user"` 
	jwt.StandardClaims
}

