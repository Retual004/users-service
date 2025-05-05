package user

import (
	"gorm.io/gorm"
	"github.com/Retual004/task-service/"  // путь какой писать?
)


type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
	GetTasksForUser(userID uint) ([]task.Task, error)
	GetUserByID(id uint) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var existingUser User
	result := r.db.First(&existingUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	r.db.Save(&existingUser)
	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var user User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return r.db.Delete(&user).Error
}

func (r *userRepository) GetTasksForUser(userID uint) ([]task.Task, error) {
    var tasks []task.Task
    if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
        return nil, err
    }
    return tasks, nil
}


// и реализация
func (r *userRepository) GetUserByID(id uint) (User, error) {
    var u User
    if err := r.db.First(&u, id).Error; err != nil {
        return User{}, err
    }
    return u, nil
}
