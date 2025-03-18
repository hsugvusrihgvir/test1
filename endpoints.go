package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// для получения всех блюд
func getDishes(c *gin.Context) {
	var dishes []Menu
	db.Find(&dishes)
	c.JSON(http.StatusOK, dishes)
}

// для получения блюда по id
func getDishByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid dish ID"})
		return
	}
	var dish Menu
	if err := db.First(&dish, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "dish not found"})
		return
	}
	c.JSON(http.StatusOK, dish)
}

// для создания нового блюда
func createDish(c *gin.Context) {
	var newDish Menu
	if err := c.BindJSON(&newDish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	db.Create(&newDish)
	c.JSON(http.StatusCreated, newDish)
}

// для обновления информации о блюде
func updateDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid dish ID"})
		return
	}

	var dish Menu
	if err := db.First(&dish, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "dish not found"})
		return
	}

	if err := c.BindJSON(&dish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	db.Save(&dish)
	c.JSON(http.StatusOK, dish)
}

// для удаления блюда
func deleteDish(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid dish ID"})
		return
	}

	var dish Menu
	if err := db.First(&dish, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "dish not found"})
		return
	}

	db.Delete(&dish)
	c.JSON(http.StatusOK, gin.H{"message": "dish deleted"})
}
