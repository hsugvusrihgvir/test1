package main

// структура для хранения учетных данных пользователя
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var users = []Credentials{
	{"user1", "password1", "admin"},
	{"user2", "password2", "user"},
}

type Menu struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float32 `json:"cost"`
	Weight      float32 `json:"weight"`
	Photo       string  `json:"photo"`
	Status      string  `json:"status"`
}

// тобы не создавалась таблица menus а открывалась menu
func (Menu) TableName() string {
	return "menu"
}
