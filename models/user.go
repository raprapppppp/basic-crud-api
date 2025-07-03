package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstname" gorm:"size:50"`
	LastName  string `json:"lastname" gorm:"size:50"`
	Age       uint   `json:"age" gorm:"not null"`
	Email     string `email:"email" gorm:"uniqueIndex;not null"`
}
