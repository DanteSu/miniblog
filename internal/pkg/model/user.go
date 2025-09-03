package model

import "time"

type UserM struct {
	CreatedAt time.Time `gorm:"column:createdAt"`      //
	Email     string    `gorm:"column:email"`          //
	ID        int64     `gorm:"column:id;primary_key"` //
	Nickname  string    `gorm:"column:nickname"`       //
	Password  string    `gorm:"column:password"`       //
	Phone     string    `gorm:"column:phone"`          //
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //
	Username  string    `gorm:"column:username"`       //
}

// TableName sets the insert table name for this struct type
func (u *UserM) TableName() string {
	return "user"
}
