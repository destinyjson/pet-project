package main

import "gorm.io/gorm"

type RequestBody struct {
	gorm.Model
	Message string `json:"message"`
}
