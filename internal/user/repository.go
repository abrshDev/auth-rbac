package user

import (
	"github.com/abrshDev/auth-rbac/pkg/utils"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) CreateUserWithPassword(username, email, password, role string) (*User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err // let handler handle response
	}
	if role == "" {
		role = "user"
	}
	user := &User{
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword,
		Role:         role,
	}
	err = r.db.Create(user).Error
	return user, err

}
func (r *Repository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
func (r *Repository) GetUserByID(id uint) (*User, error) {
	var user User

	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
