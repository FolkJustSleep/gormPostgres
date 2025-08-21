package model

type User struct {
	BaseModel
	Name	string `json:"name"`
	Email	string `json:"email" gorm:"unique"`
	Password	string `json:"password" gorm:"not null"`
	Role	string `json:"role" gorm:"default:'user'"`
}
