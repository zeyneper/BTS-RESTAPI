package models

type User struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint   `json:"age"`
	Email   string `json:"email"`
	Books   []Book `json:"books"`
}
