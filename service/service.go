package service

import (
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/shivamks5/userserv/errs"
	"github.com/shivamks5/userserv/model"
)

type UserService interface {
	GetUser(string) (model.User, error)
	CreateUser(model.User) (model.User, error)
	UpdateUser(model.User) (model.User, error)
	PatchUser(map[string]interface{}) (model.User, error)
	DeleteUser(string) error
	ListUsers(string, int, int) []model.User
}

type userService struct {
	users map[string]model.User
	mu    sync.RWMutex
}

func NewUserService() UserService {
	return &userService{
		users: make(map[string]model.User),
	}
}

func (s *userService) GetUser(id string) (model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.users[id]
	if !ok {
		return model.User{}, errs.ErrNotFound
	}
	return user, nil
}

func (s *userService) CreateUser(user model.User) (model.User, error) {
	if err := errs.ValidateUser(user); err != nil {
		return model.User{}, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.NewString()
	user.ID = id
	s.users[id] = user
	return user, nil
}

func (s *userService) UpdateUser(updatedUser model.User) (model.User, error) {
	if err := errs.ValidateUser(updatedUser); err != nil {
		return model.User{}, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	id := updatedUser.ID
	_, ok := s.users[id]
	if !ok {
		return model.User{}, errs.ErrNotFound
	}
	updatedUser.ID = id
	s.users[id] = updatedUser
	return updatedUser, nil
}

func (s *userService) PatchUser(patchedUser map[string]interface{}) (model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := patchedUser["id"].(string)
	user, ok := s.users[id]
	if !ok {
		return model.User{}, errs.ErrNotFound
	}
	if name, ok := patchedUser["name"].(string); ok {
		if !errs.CheckName(name) {
			return model.User{}, errs.ErrInvalidField
		}
		user.Name = name
	}
	if email, ok := patchedUser["email"].(string); ok {
		if !errs.CheckEmail(email) {
			return model.User{}, errs.ErrInvalidField
		}
		user.Email = email
	}
	if age, ok := patchedUser["age"].(float64); ok {
		if !errs.CheckAge(int(age)) {
			return model.User{}, errs.ErrInvalidField
		}
		user.Age = int(age)
	}
	s.users[id] = user
	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.users[id]
	if !ok {
		return errs.ErrNotFound
	}
	delete(s.users, id)
	return nil
}

func (s *userService) ListUsers(name string, minAge, maxAge int) []model.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var users []model.User = []model.User{}
	for _, user := range s.users {
		if name != "" && !strings.EqualFold(name, user.Name) {
			continue
		}
		if minAge != 0 && user.Age < minAge {
			continue
		}
		if maxAge != 0 && user.Age > maxAge {
			continue
		}
		users = append(users, user)
	}
	return users
}
