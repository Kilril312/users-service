package user

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetbyID(id uint) (*User, error) {
	var user User

	err := r.db.First(&user, id).Error

	return &user, err
}

func (r *Repository) Update(user *User) error {
	result := r.db.Model(user).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"email":    user.Email,
			"password": user.Password,
		})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *Repository) List() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}
