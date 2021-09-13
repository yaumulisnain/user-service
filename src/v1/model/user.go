package model

import "time"

type User struct {
	ID        int32      `json:"id" gorm:"primary_key;AUTO_INCREMENT;unique_index"`
	Email     string     `json:"email" validate:"min=3,max=50,alphanum,required" gorm:"not_null;unique_index"`
	Password  string     `json:"password" validate:"min=5,required" gorm:"not_null"`
	Address   string     `json:"address,omitempty" gorm:"default:null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func (User) TableName() string {
	return "user"
}

type UserLogin struct {
	Email    string `json:"email" validate:"min=3,max=50,required"`
	Password string `json:"password" validate:"min=5,required"`
}
