package repository

import (
	"clean_arch/entity"

	// "github.com/virhanali/user-management/domain/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindById(id int) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(id int) error
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

func (repo UserRepository) FindAll() ([]entity.User, error) {
	user := []entity.User{}
	if err := repo.db.Debug().Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo UserRepository) FindById(id int) (entity.User, error) {
	user := entity.User{}
	if err := repo.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo UserRepository) Update(user entity.User) (entity.User, error) {
	if err := repo.db.Debug().Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo UserRepository) Delete(id int) error {
	if err := repo.db.Debug().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return err
	}
	return nil
}
