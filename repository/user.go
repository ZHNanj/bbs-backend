package repository

import (
	"bbs-backend/model"
	"gorm.io/gorm"
)

// UserRepository 在这里定义接口，相当于Spring中的Service接口
type UserRepository interface {
	Create(user *model.User) (*model.User, error)
}

// 定义参数
type userRepository struct {
	db *gorm.DB
}

// Create 实现用户注册
func (r *userRepository) Create(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
