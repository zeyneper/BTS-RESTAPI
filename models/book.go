package models

type Book struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Title      string `json:"title"`
	PageNumber uint64 `json:"pageNumber"`
	UserID     uint   `json:"user_id"`
	User       User   `json:"user"`
}
