package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstname" gorm:"size:50"`
	LastName  string `json:"lastname" gorm:"size:50"`
	Age       uint   `json:"age" gorm:"not null"`
	Email     string `email:"email" gorm:"uniqueIndex;not null"`
}

type Users struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Firstname   string `json:"firstname" gorm:"size:50"`
	Lastname    string `json:"lastname" gorm:"size:50"`
	Email       string `json:"email" gorm:"uniqueIndex;not null"`
	Age         uint   `json:"age" gorm:"not null"`
	Birth       string `json:"birthdate"`
	Phonenumber string `json:"phoneNumber"`
}

type Account struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"type:text;default:'user'';not null"`
}

type Task struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	AccountId uint   `gorm:"not null"`
	Task      string `gorm:"not null"`
	Completed bool   `gorm:"default:false;not null"`
}