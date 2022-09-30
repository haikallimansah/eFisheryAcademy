package repository

import (
	"clean_arch/entity"

	// "github.com/virhanali/user-management/domain/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(user entity.User) (entity.User, error)
	// Store(user models.User) (models.User, error)
	// FindAll() ([]models.User, error)
	// FindById(id int) (models.User, error)
	// Update(user models.User) (models.User, error)
	// Delete(id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) Store(user entity.User) (entity.User, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
