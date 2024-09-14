package service

import "gorm.io/gorm"

type RequestBody struct {
	gorm.Model
	Message string `json:"message"`
}
