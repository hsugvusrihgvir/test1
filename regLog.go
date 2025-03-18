package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
			token, err := GenerateToken(creds.Username, user.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
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
