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

type Ingredient struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Ingredient string  `json:"ingredient"`
	Weight     float32 `json:"weight"`
	Idprovider uint    `json:"idprovider"`
}

type Passport struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Number      string `json:"number"`
	Series      string `json:"series"`
	Year        int    `json:"year"`
	WhereIssued string `json:"where_issued"`
}

type Post struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Title  string  `json:"title"`
	Salary float32 `json:"salary"`
}

type Admin struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

type Employee struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `json:"name"`
	Lastname   string  `json:"lastname"`
	Patronymic *string `json:"patronymic"` // Может быть NULL
	Email      string  `gorm:"unique" json:"email"`
	Phone      string  `gorm:"unique" json:"phone"`
	IDPassport uint    `gorm:"not null" json:"id_passport"`
	IDPost     *uint   `json:"id_post"`
	IDAdmin    *uint   `json:"id_admin"`
	Birthday   string  `json:"birthday"`
	Address    string  `json:"address"`
}

type UserStatus struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Status   string  `json:"status"`
	Discount float32 `json:"discount"`
}

type Card struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Number       string `gorm:"unique" json:"number"`
	IDUserStatus uint   `json:"id_user_status"`
}

type Customer struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Phone    string `gorm:"unique" json:"phone"`
	Email    string `gorm:"unique" json:"email"`
	Address  string `json:"address"`
	Birthday string `json:"birthday"`
}

type UserSyst struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Username   string `gorm:"unique" json:"username"`
	Password   string `json:"password"`
	IDCard     uint   `json:"id_card"`
	IDCustomer uint   `json:"id_customer"`
}

type BaseOrder struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	DateOrder  string  `json:"date_order"`
	TimeOrder  string  `json:"time_order"`
	IDEmployee *uint   `json:"id_employee"`
	Cost       float32 `json:"cost"`
	IDUser     uint    `json:"id_user"`
}

type DeliveryOrder struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Address     string `json:"address"`
	IDCourier   *uint  `json:"id_courier"`
	IDBaseOrder uint   `json:"id_base_order"`
}

type Tips struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Sum        float32 `json:"sum"`
	IDEmployee uint    `json:"id_employee"`
	IDOrder    uint    `json:"id_order"`
}

type DishOrder struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	IDDish  uint `json:"id_dish"`
	IDOrder uint `json:"id_order"`
}

type IngredientDish struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	IDIngredient uint `json:"id_ingredient"`
	IDDish       uint `json:"id_dish"`
}

type Provider struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	Company          string `json:"company"`
	Email            string `gorm:"unique" json:"email"`
	Phone            string `gorm:"unique" json:"phone"`
	IDRepresentative *uint  `json:"id_representative"`
}

type Representative struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `json:"name"`
	Lastname   string  `json:"lastname"`
	Patronymic *string `json:"patronymic"` // Может быть NULL
	Phone      string  `gorm:"unique" json:"phone"`
	Email      string  `gorm:"unique" json:"email"`
}

// тобы не создавалась таблица menus а открывалась menu
func (Menu) TableName() string {
	return "menu"
}
