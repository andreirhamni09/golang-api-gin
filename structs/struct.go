package structs

import (	
	jwt "github.com/dgrijalva/jwt-go"
)

type Users struct {
	ID       int    `json:"id"`
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
	Phone    string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	Role     string `form:"role" json:"role" xml:"role"  binding:"required"`
	Status   bool   `form:"status" json:"status" xml:"status"  binding:"required"`
}

type AuthCustomClaims struct {
	jwt.StandardClaims
}

type UsersLogin struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type Products struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      bool   `json:"status" gorm:"default:true" validate:"required"`
}

type Results struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
