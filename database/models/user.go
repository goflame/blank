package models

import "time"

type User/*Model*/ core.Model{
	Id uint `json:"id" gorm:"id"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	RegisteredAt time.Time `json:"registered_at" gorm:"registered_at"`
}
