package user //karena file dalam user, maka package user

import (
	"time"
)

type User struct {
	ID             int `gorm:"primaryKey"`
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
