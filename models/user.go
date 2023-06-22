package models

import "time"

type CreateUserRequest struct {
	FullName    string    `json:"full_name" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Gender      string    `json:"gender" binding:"required,oneof=F M"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	CountryCode string    `json:"country_code" binding:"required"`
}
