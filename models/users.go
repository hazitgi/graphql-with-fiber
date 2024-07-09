package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	FullName    string `gorm:"not null" json:"full_name"`
	CompanyName string `gorm:"not null" json:"company_name"`
	CountryID   string `gorm:"not null" json:"country_id"`
	StateID     string `gorm:"not null" json:"state_id"`
	Email       string `gorm:"unique;not null" json:"email"`
	Location    string `gorm:"not null" json:"location"`
	Address     string `gorm:"not null" json:"address"`
	Password    string `gorm:" json:"-"`
}

func NewUser() *User {
	return &User{}
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

// }
