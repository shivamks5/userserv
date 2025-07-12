package service

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/shivamks5/userserv/model"
)

var (
	ErrNotFound     = errors.New("user not found")
	ErrInvalidField = errors.New("invalid data field")
	ErrBadRequest   = errors.New("invalid request")
)

type UserService interface {
	GetUser(string) (model.User, error)
	CreateUser(model.User) (string, error)
	DeleteUser(string) error
	ListUsers() []model.User
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
		return model.User{}, ErrNotFound
	}
	return user, nil
}

func (s *userService) CreateUser(user model.User) (string, error) {
	if user.Name == "" || user.Email == "" {
		return "", ErrInvalidField
	}
	if user.Age <= 0 {
		return "", ErrInvalidField
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.NewString()
	user.ID = id
	s.users[id] = user
	return id, nil
}

func (s *userService) DeleteUser(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.users[id]
	if !ok {
		return ErrNotFound
	}
	delete(s.users, id)
	return nil
}

func (s *userService) ListUsers() []model.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var users []model.User = []model.User{}
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}
