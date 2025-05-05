package user


import (
	"github.com/Retual004/task-service/"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return s.repo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}

func (s *UserService)GetTasksForUser(userID uint) ([]task.Task, error) {
    return s.repo.GetTasksForUser(userID)
}

func (s *UserService) GetUserByID(id uint) (User, error) {
    return s.repo.GetUserByID(id)
}
