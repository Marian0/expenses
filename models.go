package main

import (
	"time"
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Name		string	`json:"name"`
	Lastname	string	`json:"last_name"`
	Email		string	`json:"email"`
	FirebaseID	string	`json:"firebase_id"`
}

type expense struct {
	gorm.Model
	Date    time.Time	`json:"date"`
	Details string		`json:"details"`
	Amount  int			`json:"amount"`
	UserID  int			`json:"user_id"`
	User    user		`gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}