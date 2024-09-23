package userService

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id int, user User) (User, error)
	DeleteUserByID(id int) (User, error)
}

type UsrRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UsrRepository {
	return &UsrRepository{db: db}
}

func (r *UsrRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (r *UsrRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, err
}

func (r *UsrRepository) UpdateUserByID(id int, user User) (User, error) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", id).Updates(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (r *UsrRepository) DeleteUserByID(id int) (User, error) {
	var user User
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
