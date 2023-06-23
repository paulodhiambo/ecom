package models

import "time"

type CreateUserRequest struct {
	FullName    string    `json:"full_name" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Gender      string    `json:"gender" binding:"required,oneof=F M"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	CountryCode string    `json:"country_code" binding:"required"`
}

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ListUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
