package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" gorm:"size:50"`
	LastName  string `json:"last_name" gorm:"size:50"`
	Age       uint   `json:"age" gorm:"not null"`
	Email     string `email:"email" gorm:"uniqueIndex;not null"`
}

type Login struct {
	ID       uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string
	Password string
}
