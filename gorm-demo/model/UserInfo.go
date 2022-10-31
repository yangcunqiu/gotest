package model

import (
	"fmt"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID          int
	Name        string `gorm:"type:varchar(255)"`
	Age         int
	PhoneNumber string `gorm:"type:varchar(255)"`
	gorm.Model
}

func (UserInfo) TableName() string {
	return "user_info"
}

func (u *UserInfo) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate hook ...")
	return
}
