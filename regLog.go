package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func login(c *gin.Context) {
	var creds Credentials
	//считываем json запрос, если ошибка то
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	//проверка логина и пароля
	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			accessToken, err1 := GenerateAccessToken(creds.Username, user.Role)
			refreshToken, err2 := GenerateRefreshToken(creds.Username, user.Role)
			if err1 != nil || err2 != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"accesssToken": accessToken, "refreshToken": refreshToken})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
}

func register(c *gin.Context) {
	var newUser Credentials
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// проверяем, существует ли пользователь
	for _, user := range users {
		if user.Username == newUser.Username {
			c.JSON(http.StatusConflict, gin.H{"message": "user already exists"})
			return
		}
	}

	// добавляем нового пользователя в список
	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func refreshAccessToken(c *gin.Context) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(requestBody.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid refresh token"})
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "refresh token expired"})
		return
	}

	accessToken, err := GenerateAccessToken(claims.Username, claims.Role) // Не создаём новый Refresh Token!
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not refresh token"})
		return
	}

	// Отправляем новый Access Token клиенту
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": requestBody.RefreshToken,
	})
}
