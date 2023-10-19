package repository

import (
	"github.com/hitoshi-w/go-lang/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *model.User) error {
  result := ur.db.Create(&user)

	return result.Error
}

func (ur *userRepository) GetByEmail(email string) (model.User, error) {
	user := model.User{}
	if err := ur.db.Where("email=?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
